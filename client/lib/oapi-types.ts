import type { components, paths } from "~/openapi.gen";

type RemoveNever<T> = T extends any ? { [S in keyof T]: RemoveNever<T[S]> } : never;

export type ApiPath = keyof paths;
export type ApiSchema = keyof components["schemas"];
export type HttpMethod =
    "get" |
    "put" |
    "post" |
    "delete" |
    "options" |
    "head" |
    "patch" |
    "trace";
export type ResponseCode = number;

type Get<
    T extends Record<string, any>,
    K extends (string | number)[]
> = 0 extends K["length"]
    ? T
    : K extends [infer F, ...infer R]
        ? F extends keyof T
            ? R extends (string | number)[]
                ? Get<T[F], R>
                : never
            : never
        : never;

export type ResponseBody<
    T extends ApiPath,
    M extends HttpMethod,
    C extends ResponseCode
> = Get<paths, [T, M, "responses", C, "content", "application/json"]>;

export type RequestBody<
    T extends ApiPath,
    M extends HttpMethod
> = Get<paths, [T, M, "requestBody", "content", "application/json"]>;

export type Parameters<
    T extends ApiPath,
    M extends HttpMethod
> = Get<paths, [T, M, "parameters"]>;