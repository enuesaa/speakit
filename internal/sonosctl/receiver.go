package sonosctl

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events

import (
	"fmt"
	"net/http"
)

func Serve2() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events/volume", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("reqa: %+v\n", req)
		w.Write(nil)
	})

	http.ListenAndServe(":2989", mux)
}
