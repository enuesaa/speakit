package main

func New() App {
	return App{}
}

type Record struct {
	Text string
	Meta map[string]string
}
type Voice struct {
	Record Record
	Mp3 []byte
}

type Generator interface {
	Start() error
	Generate() (Record, error)
	Close() error
}
type Transformer interface {
	Start() error
	Transform(record Record) (Voice, error)
	Close() error
}
type Speaker interface {
	Start() error
	Close() error
}

type App struct {}

func (a *App) UseGenerator(g Generator) {}
func (a *App) UseTransformer(t Transformer) {}
func (a *App) UseSpeaker(s Speaker) {}
