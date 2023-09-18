package main

import (
	"fmt"

	"github.com/enuesaa/speakit/controller"
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

	spec = appendSchema(spec, "feeds", &controller.FeedSchema{})
	spec = appendSchema(spec, "feeds-with-metadata", &controller.WithMetadata[controller.FeedSchema]{})
	spec = appendOp(spec, "/feeds", Op{
		Method: "GET",
		ResponseRef: "#/components/schemas/feeds-with-metadata",
		IsListReponse: true,
	})
	spec = appendOp(spec, "/feeds/{id}", Op {
		Method: "GET",
		ResponseRef: "#/components/schemas/feeds-with-metadata",
		PathParams: []string{"id"},
	})
	spec = appendOp(spec, "/feeds", Op {
		Method: "POST",
		RequestRef: "#/components/schemas/feeds",
	})
	spec = appendOp(spec, "/feeds/{id}", Op {
		Method: "DELETE",
		PathParams: []string{"id"},
	})

	spec = appendSchema(spec, "fetch", &controller.FeedfetchSchema{})
	spec = appendOp(spec, "/feeds/{id}/fetch", Op {
		Method: "POST",
		RequestRef: "#/components/schemas/fetch",
		PathParams: []string{"id"},
	})

	spec = appendSchema(spec, "programs-with-metadata", &controller.WithMetadata[controller.ProgramSchema]{})
	spec = appendOp(spec, "/programs", Op{
		Method: "GET",
		ResponseRef: "#/components/schemas/programs-with-metadata",
		IsListReponse: true,
	})
	spec = appendOp(spec, "/programs/{id}", Op {
		Method: "GET",
		ResponseRef: "#/components/schemas/programs-with-metadata",
		PathParams: []string{"id"},
	})
	spec = appendOp(spec, "/programs/{id}", Op {
		Method: "DELETE",
		PathParams: []string{"id"},
	})

	spec = appendSchema(spec, "convert", &controller.ConvertSchema{})
	spec = appendOp(spec, "/programs/{id}/convert", Op {
		Method: "POST",
		RequestRef: "#/components/schemas/convert",
		PathParams: []string{"id"},
	})

	writeYaml(spec)
}

func configure(spec openapi3.T) openapi3.T {
	spec.OpenAPI = "3.0.3"
	spec.Info = &openapi3.Info{
		Title: "Speakit API",
		Version: "0.1.0",
	}
	spec.AddServer(&openapi3.Server{
		URL: "http://localhost:3000/api/",
		Description: "local server",
	})
	spec.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}
	spec.Paths = openapi3.Paths {}

	return spec
}

func appendSchema(spec openapi3.T, name string, schema interface{}) openapi3.T {
	schemaRef, _ := openapi3gen.NewSchemaRefForValue(schema, nil)
	spec.Components.Schemas[name] = schemaRef
	return spec
}

type Op struct {
	Method string
	RequestRef string
	ResponseRef string
	IsListReponse bool
	PathParams []string
}
func appendOp(spec openapi3.T, path string, op Op) openapi3.T {
	if spec.Paths[path] == nil {
		spec.Paths[path] = &openapi3.PathItem{}
	}

	operation := openapi3.Operation{}
	operation.Summary = op.Method + " " + path
	operation.OperationID = strcase.ToLowerCamel(op.Method + path)
	operation.Parameters = openapi3.NewParameters()

	if op.PathParams != nil && len(op.PathParams) > 0 {
		for _, pathParamName := range op.PathParams {
			operation.Parameters = append(operation.Parameters, &openapi3.ParameterRef{
				Value: &openapi3.Parameter{
					In: "path",
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

	if op.RequestRef != "" {
		operation.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Ref: op.RequestRef,
						},
					},
				},
			},
		}
	}
	if op.ResponseRef != "" {
		responseSchema := &openapi3.SchemaRef{
			Ref: op.ResponseRef,
		}
		if op.IsListReponse {
			responseSchema = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "object",
					Properties: openapi3.Schemas{
						"items": &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: openapi3.TypeArray,
								Items: &openapi3.SchemaRef{
									Ref: op.ResponseRef,
								},
							},
						},
					},
				},
			}
		}

		description := ""
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
	spec.Paths[path].SetOperation(op.Method, &operation)

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
