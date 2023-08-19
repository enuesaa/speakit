//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/enuesaa/speakit/controller"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/go-yaml/yaml"
)

func main() {
	spec := openapi3.T{}
	spec = configure(spec)
	spec = appendPath(spec, "/feeds", "POST", &controller.FeedSchema{}, nil)
	// spec = appendPaths(spec)

	writeYaml(spec)
}

func configure(spec openapi3.T) openapi3.T {
	spec.OpenAPI = "3.0.3"
	spec.Info = &openapi3.Info{
		Title: "Speakit API",
		Version: "0.1.0",
	}
	spec.AddServer(&openapi3.Server{
		URL: "http://localhost:3000/api",
		Description: "local server",
	})
	spec.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}
	spec.Paths = openapi3.Paths {}

	return spec
}

func appendPath(spec openapi3.T, path string, method string, request interface{}, response interface{}) openapi3.T {
	if spec.Paths[path] == nil {
		spec.Paths[path] = &openapi3.PathItem{}
	}
	operation := openapi3.Operation{}
	operation.Summary = method + " " + path
	operation.OperationID = method + strings.ReplaceAll(path, "/", "-")
	if request != nil {
		requestSchemaRef, _ := openapi3gen.NewSchemaRefForValue(request, nil)
		operation.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Ref: "a",
						},
					},
				},
			},
		}
		spec.Components.Schemas["a"] = requestSchemaRef
	}
	if response != nil {
		responseSchemaRef, _ := openapi3gen.NewSchemaRefForValue(response, nil)
		operation.Responses = openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: responseSchemaRef,
						},
					},
				},
			},
		}
	}
	spec.Paths[path].SetOperation(method, &operation)

	return spec
}

func writeYaml(spec openapi3.T) {
	specYaml, err := yaml.Marshal(spec)
	if err != nil {
		fmt.Println("failed to marshal.")
		return
	}

	f, err := os.Create("openapi.yaml")
	if err != nil {
		fmt.Println("failed to create file.")
		return
	}
	defer f.Close()
	f.Write(specYaml)
}
