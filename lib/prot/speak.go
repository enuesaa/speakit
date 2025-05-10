package prot

type Speaker interface {
	Callfn
	Speak(Record) error
	CancelWait() error
}
