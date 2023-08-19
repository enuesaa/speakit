//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"
	"github.com/go-yaml/yaml"
	"github.com/iancoleman/strcase"
)

func main() {
	spec := openapi3.T{}
	spec = configure(spec)
	spec = defineOps(spec, "/feeds", OpsOption {
		List: true,
		View: true,
		Create: true,
		Schema: &controller.FeedSchema{},
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
		URL: "http://localhost:3000/api",
		Description: "local server",
	})
	spec.Components = &openapi3.Components{
		Schemas: openapi3.Schemas{},
	}
	spec.Paths = openapi3.Paths {}

	return spec
}

type OpsOption struct {
	List bool
	View bool
	Create bool
	Update bool
	Delete bool
	Schema interface{}
}

func defineOps(spec openapi3.T, path string, options OpsOption) openapi3.T {
	schemaName := strcase.ToLowerCamel(path)
	schemaRef := "#/components/schemas/" + schemaName
	spec = appendSchema(spec, schemaName, options.Schema)
	if options.List {
		spec = appendOp(spec, path, "GET", OpOption{responseSchemaRef: schemaRef, listResponse: true})
	}
	if options.View {
		spec = appendOp(spec, path + "/{id}", "GET", OpOption{responseSchemaRef: schemaRef})
	}
	if options.Create {
		spec = appendOp(spec, path, "POST", OpOption{requestSchemaRef: schemaRef})
	}
	if options.Update {
		spec = appendOp(spec, path, "PUT" + "/{id}", OpOption{requestSchemaRef: schemaRef})
	}
	if options.Delete {
		spec = appendOp(spec, path, "DELETE" + "/{id}", OpOption{})
	}
	return spec
}

func appendSchema(spec openapi3.T, name string, schema interface{}) openapi3.T {
	schemaRef, _ := openapi3gen.NewSchemaRefForValue(schema, nil)
	spec.Components.Schemas[name] = schemaRef
	return spec
}

type OpOption struct {
	requestSchemaRef string
	responseSchemaRef string
	listResponse bool
}
func appendOp(spec openapi3.T, path string, method string, option OpOption) openapi3.T {
	if spec.Paths[path] == nil {
		spec.Paths[path] = &openapi3.PathItem{}
	}
	operation := openapi3.Operation{}
	operation.Summary = method + " " + path
	operation.OperationID = strcase.ToLowerCamel(method + path)
	if option.requestSchemaRef != "" {
		operation.RequestBody = &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Ref: option.requestSchemaRef,
						},
					},
				},
			},
		}
	}
	if option.responseSchemaRef != "" {
		responseSchema := &openapi3.SchemaRef{
			Ref: option.responseSchemaRef,
		}
		if option.listResponse {
			responseSchema = &openapi3.SchemaRef{
				Value: &openapi3.Schema{
					Type: "object",
					Properties: openapi3.Schemas{
						"items": &openapi3.SchemaRef{
							Ref: option.responseSchemaRef,
						},
					},
				},
			}
		}

		operation.Responses = openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
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

	f, err := os.Create("openapi.yaml")
	if err != nil {
		fmt.Println("failed to create file.")
		return
	}
	defer f.Close()
	f.Write(specYaml)
}
