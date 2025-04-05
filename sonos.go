package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events
// curl -v http://192.168.3.25:1400/MediaRenderer/RenderingControl/Event \
// -H "callback: <http://m4.local:1234/sonos-event>" \
// -H "NT: upnp:event" \
// -H "Timeout: Second-1800" -X SUBSCRIBE

func getLocalIpAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)

	return addr.IP.To4().String(), nil
}

func listen2989() {
	ln, err := net.Listen("tcp", ":2989")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		fmt.Println("a")
		request, err := io.ReadAll(conn)
		if err != nil {
			panic(err)
		}
		fmt.Printf("req: %s\n", request)
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 0\r\nConnection: close\r\n\r\n"))
		conn.Close()
	}
}

func subscribeSonos() {
	go func ()  {
		listen2989()
	}()


	time.Sleep(2 * time.Second)

	// see https://stackoverflow.com/questions/19579409/how-to-subscribe-to-upnp-events

	sonosIP := discover()
	fmt.Println(sonosIP)
	url := fmt.Sprintf("http://%s:1400/MediaRenderer/RenderingControl/Event", sonosIP)

	req, err := http.NewRequest("SUBSCRIBE", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("callback", "<http://m4.local:2989>")
	req.Header.Set("NT", "upnp:event")
	req.Header.Set("Timeout", "Second-1800")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Printf("res: %+v\n", res)
	resbody, _ := io.ReadAll(res.Body)
	fmt.Printf("resbody: %s\n", string(resbody))

	time.Sleep(100 * time.Second)
}


var streamURL = "" // something mp3 url
var client = &http.Client{}

func controlSonos() {
	sonosIP := discover()
	fmt.Println(sonosIP)

	makeSetUriRequest(sonosIP)
	makePlayRequest(sonosIP)
}

func makeSetUriRequest(sonosIP string) {
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

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:1400/MediaRenderer/AVTransport/Control", sonosIP), strings.NewReader(body))
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI"`)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("res: %+v\n", res)

	resbody, _ := io.ReadAll(res.Body)
	fmt.Printf("resbody: %s\n", string(resbody))
}

func makePlayRequest(sonosIP string) {
	body := `
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:Play xmlns:u="urn:schemas-upnp-org:service:AVTransport:1">
      <InstanceID>0</InstanceID>
      <Speed>1</Speed>
    </u:Play>
  </s:Body>
</s:Envelope>`

	req, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:1400/MediaRenderer/AVTransport/Control", sonosIP), strings.NewReader(body))
	req.Header.Set("SOAPACTION", `"urn:schemas-upnp-org:service:AVTransport:1#Play"`)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("res: %+v\n", res)
}

func discover() string {
	ssdpAddr := "239.255.255.250:1900"

	msg := strings.Join([]string{
		"M-SEARCH * HTTP/1.1",
		"HOST: " + ssdpAddr,
		"MAN: \"ssdp:discover\"",
		"MX: 1",
		"ST: urn:schemas-upnp-org:device:ZonePlayer:1",
		"", "",
	}, "\r\n")

	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp4", ssdpAddr)
	if err != nil {
		panic(err)
	}

	_, err = conn.WriteTo([]byte(msg), dst)
	if err != nil {
		panic(err)
	}
	conn.SetDeadline(time.Now().Add(2 * time.Second))

	buf := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}
		res := string(buf[:n])
	
		if strings.Contains(res, "Sonos") {
			return strings.Split(addr.String(), ":")[0]
		}
	}
	panic(fmt.Errorf("err"))
}
