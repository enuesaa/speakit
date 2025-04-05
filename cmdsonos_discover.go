package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func Discover() (string, error) {
	msg := strings.Join([]string{
		"M-SEARCH * HTTP/1.1",
		"HOST: 239.255.255.250:1900",
		"MAN: \"ssdp:discover\"",
		"MX: 1",
		"ST: urn:schemas-upnp-org:device:ZonePlayer:1",
		"", "",
	}, "\r\n")

	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	dst, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if err != nil {
		return "", err
	}

	_, err = conn.WriteTo([]byte(msg), dst)
	if err != nil {
		return "", err
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
			return strings.Split(addr.String(), ":")[0], nil
		}
	}
	return "", fmt.Errorf("not found")
}
