package sonosctl

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/speakit/internal/aictl"
	"github.com/gorilla/mux"
)

func Serve() {
	router := mux.NewRouter()

	router.HandleFunc("/storage/{filename}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filename := vars["filename"]
		fmt.Println(filename)

		buf, ok := aictl.Data[filename]
		if ok {
			w.Write(buf)
			return
		}
		w.Write(nil)
	})

	router.HandleFunc("/events/volume", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("reqa: %+v\n", req)
		w.Write(nil)
	})

	http.ListenAndServe(":2989", router)
}
