package prot

type Speaker interface {
	StartUp(logger Logger) error
	Speak(Record) error
	CancelWait() error
	Close() error
}
