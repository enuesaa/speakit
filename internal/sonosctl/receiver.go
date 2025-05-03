package sonosctl

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events

import (
	"fmt"
	"io"
	"net/http"

	"github.com/enuesaa/speakit/internal/aictl"
	"github.com/gorilla/mux"
)

func (s *Sonos) StartReceiver() {
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
		fmt.Printf("req: %+v\n", req)
		w.Write(nil)
	})

	router.HandleFunc("/events/media", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("req: %+v\n", req)
		reqbody, _ := io.ReadAll(req.Body)

		fmt.Printf("reqbody: %s\n", string(reqbody))

		if s.OnMediaControl != nil {
			s.OnMediaControl()
		}
		w.Write(nil)
	})

	http.ListenAndServe(":3000", router)
}

func (s *Sonos) GetReceiverHost() string {
	return fmt.Sprintf("%s:3000", s.localIpAddr)
}
