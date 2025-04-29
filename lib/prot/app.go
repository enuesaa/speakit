package prot

type Record struct {
	Text string
	Voice []byte
	Meta map[string]string
}

type App struct {
	generator Generator
	transformers []Transformer
}

func (a *App) Transform(transformer Transformer) {
	a.transformers = append(a.transformers, transformer)
}

func (a *App) Speak(speaker Speaker) error {
	return nil
}
