package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
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

func (s *Sonos) getLocalIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)

	return addr.IP.To4().String(), nil
}

func (s *Sonos) makeSubscribe(endpoint string) (*http.Request, error) {
	url := fmt.Sprintf("http://%s:1400%s", s.ipAddr, endpoint)
	return http.NewRequest("SUBSCRIBE", url, nil)
}

type Envelope struct {
	XMLName       xml.Name `xml:"s:Envelope"`
	XmlnsS        string   `xml:"xmlns:s,attr"`
	EncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body          Body      `xml:"s:Body"`
}

type Body struct {
	Action any `xml:",any"`
}

func (s *Sonos) makePost(endpoint string, body any) (*http.Request, error) {
	url := fmt.Sprintf("http://%s:1400%s", s.ipAddr, endpoint)
	envelope := Envelope{
		XmlnsS:        "http://schemas.xmlsoap.org/soap/envelope/",
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: Body{
			Action: body,
		},
	}
	envelopbytes, err := xml.Marshal(envelope)
	if err != nil {
		return nil, err
	}
	return http.NewRequest("POST", url, bytes.NewBuffer(envelopbytes))
}

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

type SetAVTransportURI struct {
	XMLName            xml.Name `xml:"u:SetAVTransportURI"`
	XmlnsU             string   `xml:"xmlns:u,attr"`
	InstanceID         int      `xml:"InstanceID"`
	CurrentURI         string   `xml:"CurrentURI"`
	CurrentURIMetaData string   `xml:"CurrentURIMetaData"`
}

func (s *Sonos) SetUriRequest(url string) (*http.Response, error) {
	body := SetAVTransportURI{
		XmlnsU:             "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID:         0,
		CurrentURI:         url,
		CurrentURIMetaData: "",
	}
	req, err := s.makePost("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI"`)

	return s.clinet.Do(req)
}

type Play struct {
	XMLName    xml.Name `xml:"u:Play"`
	XmlnsU     string   `xml:"xmlns:u,attr"`
	InstanceID int      `xml:"InstanceID"`
	Speed      string   `xml:"Speed"`
}

func (s *Sonos) PlayRequest() (*http.Response, error) {
	body := Play{
		XmlnsU:     "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID: 0,
		Speed:      "1",
	}
	req, err := s.makePost("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#Play"`)

	return s.clinet.Do(req)
}
