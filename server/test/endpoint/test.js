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
                if (status !== "200" || status !== "201") {
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
    
    return requests;
}

function getComponent(schemas, schema) {
    const name = schema.$ref.slice('#/components/schemas/'.length);
    const component = schemas[name].properties;
    
    return Object.entries(component).reduce((acc, [key, value]) => {
        acc[key] = value.example;
        return acc;
    }, {});
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
    
    for (const req of requests) {
        try {
            const options = {
                method: req.method,
                headers: { 'Content-Type': 'application/json' },
                body: req.data ? JSON.stringify(req.data) : undefined
            };
            const response = await fetch(req.url, options);
            const data = await response.json();
            
            results += `### ${req.method} ${req.url} ${response.status} ${response.statusText}\n`;
            results += "```json\n";
            results += `response: ${JSON.stringify(data, null, "\t")}\nexpected: ${JSON.stringify(req.example, null, "\t")}\n\n`;
            results += "```\n";
            
            const diff = Diff.diffJson(data, req.example);
            const real = diff.filter(item => item.removed);
            const expected = diff.filter(item => item.added);
            diffs += `### ${req.method} ${req.url} diff\n`;
            diffs += "```json\n";
            diffs += `real: ${JSON.stringify(real, null, "\t")}\n\nexpected: ${JSON.stringify(expected, null, "\t")}\n\n`;
            diffs += "```\n";
        } catch (error) {
            results += `### ${req.method} ${req.url}\nERROR: ${error.message}\n\n`;
        }
    }
    
    fs.writeFileSync('results.md', diffs + results);
}

try {
    const doc = yaml.load(fs.readFileSync('../../../docs/openapi/gen/openapi.yaml', 'utf8'));
    const requests = getRequests(doc);
    console.log(requests);
    // await executeRequests(requests);
} catch (e) {
    console.log(e);
}