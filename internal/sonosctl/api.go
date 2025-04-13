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

func (s *Sonos) SetUri(url string) (*http.Response, error) {
	body := SetAVTransportURI{
		XmlnsU:             "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID:         0,
		CurrentURI:         url,
		CurrentURIMetaData: "",
	}
	req, err := s.post("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI"`)

	return s.clinet.Do(req)
}


type SetNextAVTransportURI struct {
	XMLName         xml.Name `xml:"u:SetNextAVTransportURI"`
	XmlnsU          string   `xml:"xmlns:u,attr"`
	InstanceID      int      `xml:"InstanceID"`
	NextURI         string   `xml:"NextURI"`
	NextURIMetaData string   `xml:"NextURIMetaData"`
}

func (s *Sonos) SetNextURI(url string) (*http.Response, error) {
	body := SetNextAVTransportURI{
		XmlnsU:          "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID:      0,
		NextURI:         url,
		NextURIMetaData: "",
	}
	req, err := s.post("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#SetNextAVTransportURI"`)

	return s.clinet.Do(req)
}

type AddURIToQueue struct {
	XMLName                        xml.Name `xml:"u:AddURIToQueue"`
	XmlnsU                         string   `xml:"xmlns:u,attr"`
	InstanceID                     string   `xml:"InstanceID"`
	EnqueuedURI                    string   `xml:"EnqueuedURI"`
	EnqueuedURIMetaData           string   `xml:"EnqueuedURIMetaData"`
	DesiredFirstTrackNumberEnqueued int `xml:"DesiredFirstTrackNumberEnqueued"`
	EnqueueAsNext                 string   `xml:"EnqueueAsNext"`
}

func (s *Sonos) AddURIToQueue(url string) (*http.Response, error) {
	body := AddURIToQueue{
		XmlnsU: "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID: "0",
		EnqueuedURI: url,
		EnqueuedURIMetaData: "",
		DesiredFirstTrackNumberEnqueued: 0,
		EnqueueAsNext: "1",
	}
	req, err := s.post("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#AddURIToQueue"`)

	return s.clinet.Do(req)
}

type Play struct {
	XMLName    xml.Name `xml:"u:Play"`
	XmlnsU     string   `xml:"xmlns:u,attr"`
	InstanceID int      `xml:"InstanceID"`
	Speed      string   `xml:"Speed"`
}

func (s *Sonos) Play() (*http.Response, error) {
	body := Play{
		XmlnsU:     "urn:schemas-upnp-org:service:AVTransport:1",
		InstanceID: 0,
		Speed:      "1",
	}
	req, err := s.post("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#Play"`)

	return s.clinet.Do(req)
}

type Next struct {
	XMLName    xml.Name `xml:"u:Next"`
	InstanceID int      `xml:"InstanceID"`
}

func (s *Sonos) Next() (*http.Response, error) {
	body := Next{
		InstanceID: 0,
	}
	req, err := s.post("/MediaRenderer/AVTransport/Control", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#Next"`)

	return s.clinet.Do(req)
}
