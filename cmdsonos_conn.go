package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func NewSonos(ipAddr string) Sonos {
	return Sonos{
		clinet: &http.Client{},
		ipAddr: ipAddr,
	}
}

type Sonos struct {
	clinet *http.Client
	ipAddr string
}

func (s *Sonos) getLocalIpAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)

	return addr.IP.To4().String(), nil
}

func (s *Sonos) subscribeSonos() (*http.Response, error) {
	localIp, err := s.getLocalIpAddress()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("http://%s:1400/MediaRenderer/RenderingControl/Event", s.ipAddr)

	req, err := http.NewRequest("SUBSCRIBE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("callback", fmt.Sprintf("<http://%s:2989>", localIp))
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	return s.clinet.Do(req)
}

func (s *Sonos) makeSetUriRequest() (*http.Response, error) {
	var streamURL = "" // something mp3 url

	body := fmt.Sprintf(`
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:SetAVTransportURI xmlns:u="urn:schemas-upnp-org:service:AVTransport:1">
      <InstanceID>0</InstanceID>
      <CurrentURI>%s</CurrentURI>
      <CurrentURIMetaData></CurrentURIMetaData>
    </u:SetAVTransportURI>
  </s:Body>
</s:Envelope>`, streamURL)

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:1400/MediaRenderer/AVTransport/Control", s.ipAddr), strings.NewReader(body))
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI"`)

	return s.clinet.Do(req)
}

func (s *Sonos) makePlayRequest() (*http.Response, error) {
	body := `
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:Play xmlns:u="urn:schemas-upnp-org:service:AVTransport:1">
      <InstanceID>0</InstanceID>
      <Speed>1</Speed>
    </u:Play>
  </s:Body>
</s:Envelope>`

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:1400/MediaRenderer/AVTransport/Control", s.ipAddr), strings.NewReader(body))
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#Play"`)

	return s.clinet.Do(req)
}
