package sonosctl

import (
	"net/http"
	"time"
)

func New() (Sonos, error) {
	sonos := Sonos{}
	sonos.clinet = &http.Client{}
	sonosIpAddr, err := sonos.discover()
	if err != nil {
		return sonos, err
	}
	sonos.ipAddr = sonosIpAddr

	localIpAddr, err := sonos.getLocalIpAddr()
	if err != nil {
		return sonos, err
	}
	sonos.localIpAddr = localIpAddr

	time.Sleep(2 * time.Second)

	return sonos, nil
}

type Sonos struct {
	clinet *http.Client
	ipAddr string
	localIpAddr string

	OnMediaControl func()
}
