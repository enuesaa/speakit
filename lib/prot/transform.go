package prot

type Transformer interface {
	StartUp(logger Logger) error
	Transform(record *Record) error
	Close() error
}
