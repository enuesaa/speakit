package prot

type Skipper interface {
	Callfn
	ShouldSkip(Record) bool
}
