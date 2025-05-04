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

func (g *BeepSpeaker) StartUp() error {
	return nil
}

func (g *BeepSpeaker) Speak(record Record) (time.Duration, error) {
	reader := bytes.NewBuffer(record.Voice)
	readcloser := io.NopCloser(reader)

	streamer, format, err := mp3.Decode(readcloser)
	if err != nil {
		return 0, err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)

	streamerreal := buffer.Streamer(0, buffer.Len())
	speaker.Play(streamerreal)

	duration := time.Duration(buffer.Len()) * time.Second / time.Duration(format.SampleRate)

	return duration, nil
}

func (g *BeepSpeaker) Stop() error {
	return nil
}

func (g *BeepSpeaker) Close() error {
	return nil
}
