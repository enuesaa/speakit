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
	stopped bool
}

func (g *BeepSpeaker) Inject(log *LogBehavior) {
	g.log = log
}

func (g *BeepSpeaker) StartUp() error {
	g.stopped = true
	return nil
}

func (g *BeepSpeaker) Speak(record Record) error {
	g.stopped = false

	reader := bytes.NewBuffer(record.Voice)
	readcloser := io.NopCloser(reader)

	streamer, format, err := mp3.Decode(readcloser)
	if err != nil {
		return err
	}
	defer streamer.Close()

	// TODO
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamerreal := buffer.Streamer(0, buffer.Len())

	g.ctrl = &beep.Ctrl{Streamer: streamerreal, Paused: false}
	withCallback := beep.Seq(g.ctrl, beep.Callback(func() {
		g.stopped = true
	}))
	speaker.Play(withCallback)

	return nil
}

func (g *BeepSpeaker) CancelWait() error {
	if g.ctrl != nil {
		speaker.Lock()
		g.ctrl.Paused = true
		speaker.Unlock()
	}
	g.stopped = true
	return nil
}

func (g *BeepSpeaker) Close() error {
	return nil
}

func (g *BeepSpeaker) IsStopped() bool {
	return g.stopped
}
