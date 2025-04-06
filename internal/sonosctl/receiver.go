package sonosctl

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events

import (
	"fmt"
	"net/http"
)

type Receiver struct{}

func (r *Receiver) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("req: %+v\n", request)
	writer.Write(nil)
}

func Serve() {
	server := &http.Server{
		Addr:    ":2989",
		Handler: &Receiver{},
	}
	server.ListenAndServe()
}
