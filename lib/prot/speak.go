package prot

type Speaker interface {
	StartUp(app *App) error
	Speak(Record) error
	Stop() error
	Close() error
}
