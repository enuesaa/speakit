package prot

func New(g Generator) *App {
	return &App{
		generator: g,
	}
}
