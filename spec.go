package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func GenerateOpenAPISpec() openapi3.T {
	spec := openapi3.T{}
	spec.OpenAPI = "3.0.3"
	spec.Info = &openapi3.Info{
		Title: "",
		Version: "",
	}
	spec.AddServer(&openapi3.Server{
		URL: "",
		Description: "",
	})
	spec.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}

	return spec
}