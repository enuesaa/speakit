//go:build dev

package admin

func init() {
	go RunDevCmd()
}

var Serve = ServeDev
