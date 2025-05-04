package prot

type Controller interface {
	StartUp(app *App) error
	Close() error
}
