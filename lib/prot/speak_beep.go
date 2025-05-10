package prot

import (
	"bytes"
	"fmt"
	"io"
	"time"
	_ "embed"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

//go:embed assets/connected.mp3
var assetConnected []byte

type BeepSpeaker struct {
	Storage map[string][]byte

	log  *LogBehavior
	ctrl *beep.Ctrl // judge is-playing with this value
}

func (g *BeepSpeaker) Inject(log *LogBehavior) {
	g.log = log
}

func (g *BeepSpeaker) StartUp() error {
	return g.play(assetConnected)
}

func (g *BeepSpeaker) Speak(record Record) error {
	for _, segment := range record.Segments {
		voice, err := record.Speech(segment)
		if err != nil {
			return err
		}
		if err := g.play(voice); err != nil {
			return err
		}
	}
	return nil
}

func (g *BeepSpeaker) play(voice []byte) error {
	reader := bytes.NewBuffer(voice)
	readcloser := io.NopCloser(reader)

	streamer, format, err := mp3.Decode(readcloser)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buf := beep.NewBuffer(format)
	buf.Append(streamer)
	bufstreamer := buf.Streamer(0, buf.Len())

	if err := g.wait(); err != nil {
		return nil
	}
	g.ctrl = &beep.Ctrl{Streamer: bufstreamer, Paused: false}
	withCallback := beep.Seq(g.ctrl, beep.Callback(func() {
		g.ctrl = nil
	}))
	speaker.Play(withCallback)

	return nil
}

func (g *BeepSpeaker) wait() error {
	for {
		if g.ctrl == nil {
			break
		}
		if g.ctrl.Paused {
			g.ctrl = nil
			return fmt.Errorf("end")
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (g *BeepSpeaker) CancelWait() error {
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

