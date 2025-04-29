package main

func New() AppGenerate {
	return &App{}
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

func (a *App) Generate(g Generator) AppTransform {
	return a
}
func (a *App) Transform(t Transformer) AppTransformSpeak {
	return a
}
func (a *App) Speak(s Speaker) *App {
	return a
}

type AppGenerate interface {
	Generate(g Generator) AppTransform
}
type AppTransform interface {
	Transform(t Transformer) AppTransformSpeak
}
type AppTransformSpeak interface {
	AppTransform
	AppSpeak
}
type AppSpeak interface {
	Speak(s Speaker)
}

func example() {
	app := New().
		Generate().
		Transform()
}
