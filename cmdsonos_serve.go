package main

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events
// curl -v http://192.168.3.25:1400/MediaRenderer/RenderingControl/Event \
// -H "callback: <http://m4.local:1234/sonos-event>" \
// -H "NT: upnp:event" \
// -H "Timeout: Second-1800" -X SUBSCRIBE

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
