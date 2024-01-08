//go:build dev

package web

func init() {
	go RunDevCmd()
}

var Serve = ServeDev