package prot

type Transformer interface {
	StartUp(app *App) error
	Transform(record *Record) error
	Close() error
}
