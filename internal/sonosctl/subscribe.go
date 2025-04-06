package sonosctl

import (
	"fmt"
	"net/http"
)

func (s *Sonos) SubscribeVolumeControl() (*http.Response, error) {
	localIp, err := s.getLocalIpAddr()
	if err != nil {
		return nil, err
	}
	req, err := s.makeSubscribe("/MediaRenderer/RenderingControl/Event")
	if err != nil {
		return nil, err
	}
	req.Header.Set("callback", fmt.Sprintf("<http://%s:2989>", localIp))
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	return s.clinet.Do(req)
}