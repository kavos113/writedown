import { defineEventHandler, proxyRequest } from 'h3';

export default defineEventHandler(async (event) => {
    const target = 'http://localhost:8080';
    const url = event.node.req.url?.replace(/^\/api/, '') || '';
    return proxyRequest(event, target + url);
});