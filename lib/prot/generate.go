package prot

type Generator interface {
	StartUp(app *App) error
	Generate() (Record, error)
	Close() error
}
