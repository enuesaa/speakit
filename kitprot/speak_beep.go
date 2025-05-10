package kitprot

import (
	"bytes"
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
	ctrl *beep.Ctrl
	stopped bool
	queue []Record
}

func (g *BeepSpeaker) Inject(log *LogBehavior) {
	g.log = log
}

func (g *BeepSpeaker) StartUp() error {
	if err := g.play(assetConnected); err != nil {
		return err
	}
	g.queue = make([]Record, 0)

	go g.process()

	return nil
}

func (g *BeepSpeaker) Speak(record Record) error {
	g.queue = append(g.queue, record)
	return nil
}

func (g *BeepSpeaker) process() {
	for {
		if len(g.queue) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		record := g.queue[0]
		if len(g.queue) > 1 {
			g.queue = g.queue[1:]
		} else {
			g.queue = make([]Record, 0)
		}
		if err := g.processRecord(record); err != nil {
			panic(err)
		}
	}
}

func (g *BeepSpeaker) processRecord(record Record) error {
	g.stopped = false

	for _, segment := range record.Segments {
		if g.stopped {
			break
		}
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
	if g.stopped {
		return nil
	}

	g.ctrl = &beep.Ctrl{Streamer: bufstreamer, Paused: false}
	withCallback := beep.Seq(g.ctrl, beep.Callback(func() {
		g.ctrl = nil
	}))
	speaker.Play(withCallback)

	return nil
}

func (g *BeepSpeaker) wait() {
	for range 100 {
		if g.ctrl == nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
	panic("wait too long")
}

func (g *BeepSpeaker) Cancel() error {
	if g.ctrl != nil {
		speaker.Lock()
		g.ctrl.Paused = true
		speaker.Unlock()
		g.stopped = true
		g.ctrl = nil
	}
	return nil
}

func (g *BeepSpeaker) Close() error {
	return nil
}

