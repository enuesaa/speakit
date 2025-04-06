package sonosctl

import (
	"encoding/xml"
	"net/http"
)

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
