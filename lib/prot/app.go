package prot

type Record struct {
	Text string
	Meta map[string]string
}
type Voice struct {
	Record Record
	Mp3 []byte
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
