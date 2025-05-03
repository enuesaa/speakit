package prot

import (
	"fmt"
	"net/http"
	"time"

	"github.com/enuesaa/speakit/lib/sonosctl"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SonosSpeaker struct {
	sonos sonosctl.Sonos

	Storage map[string][]byte
}

func (g *SonosSpeaker) StartUp() error {
	sonos, err := sonosctl.New()
	if err != nil {
		return err
	}
	g.sonos = sonos
	g.Storage = make(map[string][]byte)

	go g.serve()

	return nil
}

func (g *SonosSpeaker) serve() error {
	router := mux.NewRouter()
	router.HandleFunc("/storage/{filename}", func(w http.ResponseWriter, r *http.Request) {
		filename := mux.Vars(r)["filename"]
		fmt.Printf("requested: %s\n", filename)

		data, ok := g.Storage[filename]
		if ok {
			w.Write(data)
		}
	})
	return http.ListenAndServe(":3000", router)
}

func (g *SonosSpeaker) Next(record Record) (time.Duration, error) {
	filename := fmt.Sprintf("%s.mp3", uuid.NewString())
	g.Storage[filename] = record.Voice

	url := fmt.Sprintf("http://%s/storage/%s", g.sonos.GetReceiverHost(), filename)
	if _, err := g.sonos.SetUri(url); err != nil {
		return 0, err
	}
	if _, err := g.sonos.Play(); err != nil {
		return 0, err
	}
	return 0, nil
}

func (g *SonosSpeaker) Stop() error {
	return nil
}

func (g *SonosSpeaker) Close() error {
	return nil
}
