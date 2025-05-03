package prot

import (
	"bytes"
	"io"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type BeepSpeaker struct {
	Storage map[string][]byte
}

func (g *BeepSpeaker) Start() error {
	return nil
}

func (g *BeepSpeaker) Next(record Record) error {
	reader := bytes.NewBuffer(record.Voice)
	readcloser := io.NopCloser(reader)

	streamer, format, err := mp3.Decode(readcloser)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speedy := beep.ResampleRatio(4, 1.2, streamer)

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(speedy)

	return nil
}

func (g *BeepSpeaker) Stop() error {
	return nil
}

func (g *BeepSpeaker) Close() error {
	return nil
}
