package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func NewSonos(ipAddr string) Sonos {
	return Sonos{
		ipAddr: ipAddr,
	}
}

type Sonos struct {
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

func (s *Sonos) subscribeSonos() error {
	localIp, err := s.getLocalIpAddress()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s:1400/MediaRenderer/RenderingControl/Event", s.ipAddr)

	req, err := http.NewRequest("SUBSCRIBE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("callback", fmt.Sprintf("<http://%s:2989>", localIp))
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf("res: %+v\n", res)
	resbody, _ := io.ReadAll(res.Body)
	fmt.Printf("resbody: %s\n", string(resbody))

	return nil
}

func (s *Sonos) makeSetUriRequest() error {
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

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Printf("res: %+v\n", res)

	resbody, _ := io.ReadAll(res.Body)
	fmt.Printf("resbody: %s\n", string(resbody))
	return nil
}

func (s *Sonos) makePlayRequest() error {
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

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Printf("res: %+v\n", res)

	return nil
}
