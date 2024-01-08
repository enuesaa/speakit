package main

import (
	"flag"
)

var emitOpenapiFlag bool
var tryFetchFlag string

func init() {
	flag.BoolVar(&emitOpenapiFlag, "emit-openapi", false, "create openapi.yaml")
	flag.StringVar(&tryFetchFlag, "try-fetch", "", "try to fetch rss feed. please provide url.")
}

func main() {
	flag.Parse()

	if emitOpenapiFlag {
		emitOpenapi()
	} else if tryFetchFlag != "" {
		TryFetch(tryFetchFlag)
	} else {
		serve()
	}
}
