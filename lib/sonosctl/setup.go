package sonosctl

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var discoverMsg = `M-SEARCH * HTTP/1.1
HOST: 239.255.255.250:1900
MAN: "ssdp:discover"
MX: 1
ST: urn:schemas-upnp-org:device:ZonePlayer:1

`

func (s *Sonos) discover() (string, error) {
	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if err != nil {
		return "", err
	}
	if _, err := conn.WriteTo([]byte(discoverMsg), dst); err != nil {
		return "", err
	}
	if err := conn.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		return "", err
	}

	buf := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}
		data := string(buf[:n])

		if strings.Contains(data, "Sonos") {
			udpAddr := addr.(*net.UDPAddr)
			return udpAddr.IP.To4().String(), nil
		}
	}
	return "", fmt.Errorf("not found")
}

func (s *Sonos) getLocalIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	udpAddr := conn.LocalAddr().(*net.UDPAddr)

	return udpAddr.IP.To4().String(), nil
}
