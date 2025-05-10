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

	log     *LogBehavior
	ctrl    *beep.Ctrl
}

func (g *BeepSpeaker) Inject(log *LogBehavior) {
	g.log = log
}

func (g *BeepSpeaker) StartUp() error {
	return nil
}

func (g *BeepSpeaker) Speak(record Record) error {
	playing := false

	for {
		voice, err := record.Speech()
		if err != nil {
			return err
		}
		if voice == nil {
			return nil
		}
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

		for {
			if !playing {
				break
			}
			time.Sleep(1 * time.Second)
		}

		g.ctrl = &beep.Ctrl{Streamer: bufstreamer, Paused: false}
		withCallback := beep.Seq(g.ctrl, beep.Callback(func() {
			playing = false
		}))
		speaker.Play(withCallback)
	}
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

func (g *BeepSpeaker) IsStopped() bool {
	return false
}
