package kitprot

type Skipper interface {
	Callfn
	ShouldSkip(Record) bool
}
