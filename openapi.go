package main

import (
	"fmt"

	"github.com/enuesaa/speakit/pkg/controller"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/go-yaml/yaml"
	"github.com/iancoleman/strcase"
)

// run following command.
// go run . -emit-openapi > ./apps/admin/openapi.yaml
func emitOpenapi() {
	spec := openapi3.T{}
	spec = configure(spec)

	spec = appendSchema(spec, "empty", &struct{}{})
	spec = appendSchema(spec, "feeds", &controller.FeedSchema{})
	spec = appendSchema(spec, "feeds-with-metadata", &controller.WithMetadata[controller.FeedSchema]{})
	spec = appendOperation(spec, "GET", "/feeds", OperationConfig{
		ListResponseSchema: "feeds-with-metadata",
	})
	spec = appendOperation(spec, "GET", "/feeds/{id}", OperationConfig{
		ResponseSchema: "feeds-with-metadata",
		PathParams:     []string{"id"},
	})
	spec = appendOperation(spec, "POST", "/feeds", OperationConfig{
		RequestSchema:  "feeds",
		ResponseSchema: "empty",
	})
	spec = appendOperation(spec, "DELETE", "/feeds/{id}", OperationConfig{
		PathParams:     []string{"id"},
		ResponseSchema: "empty",
	})

	spec = appendSchema(spec, "fetch", &controller.FeedfetchSchema{})
	spec = appendOperation(spec, "POST", "/feeds/{id}/fetch", OperationConfig{
		PathParams:     []string{"id"},
		RequestSchema:  "fetch",
		ResponseSchema: "empty",
	})

	spec = appendSchema(spec, "programs-with-metadata", &controller.WithMetadata[controller.ProgramSchema]{})
	spec = appendOperation(spec, "GET", "/programs", OperationConfig{
		ListResponseSchema: "programs-with-metadata",
	})
	spec = appendOperation(spec, "GET", "/programs/{id}", OperationConfig{
		PathParams:     []string{"id"},
		ResponseSchema: "programs-with-metadata",
	})
	spec = appendOperation(spec, "DELETE", "/programs/{id}", OperationConfig{
		PathParams:     []string{"id"},
		ResponseSchema: "empty",
	})

	spec = appendSchema(spec, "convert", &controller.ConvertSchema{})
	spec = appendOperation(spec, "POST", "/programs/{id}/convert", OperationConfig{
		PathParams:     []string{"id"},
		RequestSchema:  "convert",
		ResponseSchema: "empty",
	})

	writeYaml(spec)
}

func configure(spec openapi3.T) openapi3.T {
	spec.OpenAPI = "3.0.3"
	spec.Info = &openapi3.Info{
		Title:   "Speakit API",
		Version: "0.1.0",
	}
	spec.AddServer(&openapi3.Server{
		URL:         "http://localhost:3000/api/",
		Description: "local server",
	})
	spec.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}
	spec.Paths = openapi3.Paths{}

	return spec
}

func appendSchema(spec openapi3.T, name string, schema interface{}) openapi3.T {
	schemaRef, _ := openapi3gen.NewSchemaRefForValue(schema, nil)
	spec.Components.Schemas[name] = schemaRef
	return spec
}

type OperationConfig struct {
	PathParams         []string
	RequestSchema      string
	ResponseSchema     string
	ListResponseSchema string
}

func appendOperation(spec openapi3.T, method string, path string, config OperationConfig) openapi3.T {
	if spec.Paths[path] == nil {
		spec.Paths[path] = &openapi3.PathItem{}
	}

	operation := openapi3.Operation{}
	operation.Summary = method + " " + path
	operation.OperationID = strcase.ToLowerCamel(method + path)
	operation.Parameters = openapi3.NewParameters()

	// params
	if config.PathParams != nil && len(config.PathParams) > 0 {
		for _, pathParamName := range config.PathParams {
			operation.Parameters = append(operation.Parameters, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					In:   "path",
					Name: pathParamName,
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
					Required: true,
				},
			})
		}
	}

	// request schema
	if config.RequestSchema != "" {
		operation.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Ref: "#/components/schemas/" + config.RequestSchema,
						},
					},
				},
			},
		}
	}

	// response schema
	if config.ResponseSchema != "" {
		description := ""
		responseSchema := &openapi3.SchemaRef{
			Ref: "#/components/schemas/" + config.ResponseSchema,
		}
		operation.Responses = openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &description,
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: responseSchema,
						},
					},
				},
			},
		}
	}
	if config.ListResponseSchema != "" {
		description := ""
		responseSchema := &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: "object",
				Properties: openapi3.Schemas{
					"items": &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: openapi3.TypeArray,
							Items: &openapi3.SchemaRef{
								Ref: "#/components/schemas/" + config.ListResponseSchema,
							},
						},
					},
				},
			},
		}
		operation.Responses = openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &description,
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: responseSchema,
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

	fmt.Printf("%s", specYaml)
}
