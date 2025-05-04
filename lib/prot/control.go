package prot

type Controller interface {
	StartUp(logger Logger, app *App) error
	Close() error
}
