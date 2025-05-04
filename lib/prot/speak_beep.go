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

	ctrl    *beep.Ctrl
}

func (g *BeepSpeaker) StartUp() error {
	return nil
}

func (g *BeepSpeaker) Speak(record Record) error {
	reader := bytes.NewBuffer(record.Voice)
	readcloser := io.NopCloser(reader)

	streamer, format, err := mp3.Decode(readcloser)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamerreal := buffer.Streamer(0, buffer.Len())

	g.ctrl = &beep.Ctrl{Streamer: streamerreal, Paused: false}
	speaker.Play(g.ctrl)
	g.wait(buffer, format)

	return nil
}

func (g BeepSpeaker) wait(buffer *beep.Buffer, format beep.Format) {
	duration := time.Duration(buffer.Len()) * time.Second / time.Duration(format.SampleRate)
	time.Sleep(duration)
}

func (g *BeepSpeaker) Stop() error {
	if g.ctrl != nil {
		speaker.Lock()
		g.ctrl.Paused = true
		speaker.Unlock()
	}
	return nil
}

func (g *BeepSpeaker) Close() error {
	return nil
}
