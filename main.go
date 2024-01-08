package main

import (
	"flag"
)

var emitOpenapiFlag bool

func init() {
	flag.BoolVar(&emitOpenapiFlag, "emit-openapi", false, "create openapi.yaml")
}

func main() {
	flag.Parse()

	if emitOpenapiFlag {
		emitOpenapi()
	} else {
		serve()
	}
}
