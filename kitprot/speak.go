package kitprot

type Speaker interface {
	Callfn
	Speak(Record) error
	Cancel() error
}
