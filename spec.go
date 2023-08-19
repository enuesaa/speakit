//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
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