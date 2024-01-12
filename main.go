package main

import (
	"flag"
)

var emitOpenapiFlag bool
var tryFetchFlag bool

func init() {
	flag.BoolVar(&emitOpenapiFlag, "emit-openapi", false, "create openapi.yaml")
	flag.BoolVar(&tryFetchFlag, "tryfetch", false, "try to fetch rss feed. please provide url.")
}

func main() {
	flag.Parse()

	if emitOpenapiFlag {
		emitOpenapi()
	} else if tryFetchFlag {
		TryFetch()
	} else {
		serve()
	}
}
