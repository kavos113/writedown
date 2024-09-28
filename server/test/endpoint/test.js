const yaml = require('js-yaml');
const fs = require('fs');
const Diff = require('diff');

function getRequests(doc) {
    const requests = [];
    
    for (const path in doc.paths) {
        for (const method in doc.paths[path]) {
            if (method === "delete") continue;
            
            const req = {
                method: method.toUpperCase(),
                url: path,
                data: doc.paths[path][method].requestBody ? getComponent(doc.components.schemas, doc.paths[path][method].requestBody.content["application/json"].schema) : undefined,
            };
            
            const responses = doc.paths[path][method].responses;
            for (const status in responses) {
                if (status !== "200" && status !== "201") {
                    break;
                }
                
                req.example = responses[status].content ? getComponent(doc.components.schemas, responses[status].content["application/json"].schema) : undefined;
            }
            
            const pathParams = doc.paths[path][method].parameters?.filter(param => param.in === "path");
            const queryParams = doc.paths[path][method].parameters?.filter(param => param.in === "query");
            if (pathParams) {
                for (const param of pathParams) {
                    req.url = setPathParameter(req.url, param.schema);
                }
            }
            if (queryParams) {
                req.url = setQueryParameters(req.url, queryParams);
            }
            
            if (doc.paths[path][method].tags.includes("login")) {
                requests.unshift(req);
            } else {
                requests.push(req);
            }
        }
    }
    // signupを先に
    requests.sort((a, b) => a.url.includes("signup") ? -1 : 0);
    
    requests.forEach(req => {
        req.url = `http://localhost:8080${req.url}`;
    });
    
    return requests;
}

function getComponent(schemas, schema) {
    if (schema.type && schema.type === "array") {
        schema = schema.items;
        schema.array = true;
    }
    const name = schema.$ref.slice('#/components/schemas/'.length);
    let component = schemas[name].properties;
    
    
    if (schemas[name].items?.$ref) {
        component = getComponent(schemas, schemas[name].items);
        schema.inRef = true;
    }
    
    if (schema.inRef) {
        return schema.array ? [component] : component;
    }
    
    const ret = Object.entries(component).reduce((acc, [key, value]) => {
        acc[key] = value.example;
        return acc;
    }, {});
    
    return schema.array ? [ret] : ret;
}

function setPathParameter(path, schema) {
    const pathParameter = path.match(/{\w+}/g)[0];
    if (pathParameter) {
        path = path.replace(pathParameter, schema.example);
    }
    
    return path;
}

function setQueryParameters(path, parameters) {
    for (const param of parameters) {
        path = path.concat(`?${param.name}=${param.schema.example}`);
    }
    
    return path;
}

async function executeRequests(requests) {
    let results = "## Results\n";
    let diffs = "## Diffs\n";
    
    // signup
    try {
        const options = {
            method: requests[0].method,
            headers: { 'Content-Type': 'application/json' },
            body: requests[0].data ? JSON.stringify(requests[0].data) : undefined
        };
        const response = await fetch(requests[0].url, options);
        
        results += `### ${requests[0].method} ${requests[0].url} ${response.status} ${response.statusText}\n`;
        
        console.log(`Finish: ${requests[0].method} ${requests[0].url}, ${response.status} ${response.statusText}`);
    } catch (error) {
        results += `### ${requests[0].method} ${requests[0].url}\nERROR: ${error.message}\n\n`;
        console.log(`Error: ${requests[0].method} ${requests[0].url}`);
    }
    console.log("Signup done");
    
    let session;
    
    // login
    await new Promise(resolve => setTimeout(resolve, 300));
    try {
        const options = {
            method: requests[1].method,
            headers: { 'Content-Type': 'application/json' },
            body: requests[1].data ? JSON.stringify(requests[1].data) : undefined
        };
        const response = await fetch(requests[1].url, options);
        
        session = response.headers.get('Set-Cookie');
        
        results += `### ${requests[1].method} ${requests[1].url} ${response.status} ${response.statusText}\n`;
        
        console.log(`Finish: ${requests[1].method} ${requests[1].url}, ${response.status} ${response.statusText}`);
    } catch (error) {
        results += `### ${requests[1].method} ${requests[1].url}\nERROR: ${error.message}\n\n`;
        console.log(`Error: ${requests[1].method} ${requests[1].url}`);
    }
    
    requests = requests.slice(2);
    
    for (const req of requests) {
        await new Promise(resolve => setTimeout(resolve, 300));
        try {
            const options = {
                method: req.method,
                headers: {
                    'Content-Type': 'application/json',
                    'Cookie': session
                },
                body: req.data ? JSON.stringify(req.data) : undefined
            };
            const response = await fetch(req.url, options);
            const text = await response.text();
            
            let data;
            try {
                data = JSON.parse(text);
            } catch (e) {
                results += `### ${req.method} ${req.url} ${response.status} ${response.statusText}\n`;
                results += "```\n";
                results += `${text}\n`;
                results += "```\n";
                console.log(`Finish: ${req.method} ${req.url}, ${response.status} ${response.statusText}`);
                continue;
            }
            
            results += `### ${req.method} ${req.url} ${response.status} ${response.statusText}\n`;
            results += "```json\n";
            results += `response: ${JSON.stringify(data, null, "\t")}\nexpected: ${JSON.stringify(req.example, null, "\t")}\n\n`;
            results += "```\n";
            
            const diff = Diff.diffJson(data, req.example);
            const real = diff.filter(item => item.removed);
            const expected = diff.filter(item => item.added);
            const realText = real.map(item => {
                if (item.value.includes("created_at") || item.value.includes("updated_at")) {
                    return "";
                }
                return "+" + item.value;
            });
            const expectedText = expected.map(item => {
                if (item.value.includes("created_at") || item.value.includes("updated_at")) {
                    return "";
                }
                return "-" + item.value;
            });
            let diffText = "";
            for (let i = 0; i < realText.length; i++) {
                diffText += `${realText[i]}${expectedText[i]}`;
            }
            if (diffText === "") {
                console.log(`Finish: ${req.method} ${req.url}, ${response.status} ${response.statusText}`);
                continue;
            }
            diffs += `### ${req.method} ${req.url} diff\n`;
            diffs += "```json\n";
            diffs += `${diffText}\n`;
            diffs += "```\n\n";
            
            console.log(`Finish: ${req.method} ${req.url}, ${response.status} ${response.statusText}`);
        } catch (error) {
            results += `### ${req.method} ${req.url}\nERROR: ${error.message}\n\n`;
            console.log(`Error: ${req.method} ${req.url}`);
        }
    }
    
    fs.writeFileSync('results.md', diffs + results);
}

try {
    const doc = yaml.load(fs.readFileSync('../../../docs/openapi/gen/openapi.yaml', 'utf8'));
    const requests = getRequests(doc);
    //console.dir(requests, { depth: null });
    executeRequests(requests).then(() => console.log("done"));
} catch (e) {
    console.log(e);
}