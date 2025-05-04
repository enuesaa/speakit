package prot

type Transformer interface {
	StartUp() error
	Transform(record *Record) error
	Close() error
}
