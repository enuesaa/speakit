package sonosctl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

func (s *Sonos) calcUrl(endpoint string) string {
	return fmt.Sprintf("http://%s:1400%s", s.ipAddr, endpoint)
}

func (s *Sonos) subscribe(endpoint string) (*http.Request, error) {
	return http.NewRequest("SUBSCRIBE", s.calcUrl(endpoint), nil)
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

func (s *Sonos) post(endpoint string, body any) (*http.Request, error) {
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
	return http.NewRequest("POST", s.calcUrl(endpoint), bytes.NewBuffer(envelopbytes))
}
