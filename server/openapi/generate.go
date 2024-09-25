package openapi

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./config/model-cfg.yaml ../../docs/openapi/gen/openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./config/server-cfg.yaml ../../docs/openapi/gen/openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./config/spec-cfg.yaml ../../docs/openapi/gen/openapi.yaml
