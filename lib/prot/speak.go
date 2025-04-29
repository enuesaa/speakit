package prot

import "github.com/enuesaa/speakit/lib/sonosctl"

type Speaker interface {
	Start() error
	Next(Record) error
	Stop() error
	Close() error
}

type SonosSpeaker struct {
	sonos sonosctl.Sonos
}

func (g *SonosSpeaker) Start() error {
	sonos, err := sonosctl.New()
	if err != nil {
		return err
	}
	g.sonos = sonos

	// start web server

	return nil
}

func (g *SonosSpeaker) Next(record Record) error {
	return nil
}

func (g *SonosSpeaker) Stop() error {
	return nil
}

func (g *SonosSpeaker) Close() error {
	return nil
}
