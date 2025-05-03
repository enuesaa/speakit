package sonosctl

import (
	"fmt"
	"net/http"
)

func (s *Sonos) SubscribeVolumeControl() (*http.Response, error) {
	req, err := s.subscribe("/MediaRenderer/RenderingControl/Event")
	if err != nil {
		return nil, err
	}
	req.Header.Set("callback", fmt.Sprintf("<http://%s:2989/events/volume>", s.localIpAddr))
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	return s.clinet.Do(req)
}

func (s *Sonos) SubscribeMediaControl() (*http.Response, error) {
	req, err := s.subscribe("/MediaRenderer/AVTransport/Event")
	if err != nil {
		return nil, err
	}
	req.Header.Set("callback", fmt.Sprintf("<http://%s:2989/events/media>", s.localIpAddr))
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	return s.clinet.Do(req)
}
