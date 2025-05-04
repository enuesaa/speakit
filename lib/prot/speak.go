package prot

type Speaker interface {
	StartUp() error
	Speak(Record) error
	Stop() error
	Close() error
}
