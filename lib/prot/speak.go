package prot

type Speaker interface {
	StartUp(logger Logger) error
	Speak(Record) error
	Stop() error
	Close() error
}
