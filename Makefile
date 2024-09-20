oapi:
	npx @redocly/cli bundle -o docs/openapi/gen --ext yaml docs/openapi/paths/*.yaml docs/openapi/schemas/*.yaml docs/openapi/openapi.yaml
	go generate server/openapi/generate.go
	cd client && npm run oapi