//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/go-yaml/yaml"
)

func main() {
	spec := GenerateOpenAPISpec()

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

func GenerateOpenAPISpec() openapi3.T {
	spec := openapi3.T{}
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

	feedSchemaRef, _ := openapi3gen.NewSchemaRefForValue(&controller.FeedSchema{}, nil)

	spec.Paths = openapi3.Paths {
		"/feeds": &openapi3.PathItem{
			Post: &openapi3.Operation{
				Summary: "Add feed",
				OperationID: "postFeeds",
				RequestBody: &openapi3.RequestBodyRef{
					Value: &openapi3.RequestBody{
						Content: openapi3.Content{
							"application/json": &openapi3.MediaType{
								Schema: feedSchemaRef,
							},
						},
					},
				},
			},
		},
	}

	return spec
}