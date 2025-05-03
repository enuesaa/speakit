package prot

func GenerateFrom(g Generator) *App {
	return &App{
		generator: g,
	}
}
