package prot

type Skipper interface {
	StartUp(logger Logger) error
	ShouldSkip(Record) bool
	Close() error
}
