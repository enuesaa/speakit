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

	log  *LogBehavior
	ctrl *beep.Ctrl
	playing bool
}

func (g *BeepSpeaker) Inject(log *LogBehavior) {
	g.log = log
}

func (g *BeepSpeaker) StartUp() error {
	g.playing = false
	return nil
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

	g.wait()

	g.ctrl = &beep.Ctrl{Streamer: bufstreamer, Paused: false}
	withCallback := beep.Seq(g.ctrl, beep.Callback(func() {
		g.playing = false
	}))
	speaker.Play(withCallback)

	return nil
}

func (g *BeepSpeaker) wait() {
	for {
		if !g.playing {
			break
		}
		time.Sleep(1 * time.Second)
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

