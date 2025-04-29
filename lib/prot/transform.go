package prot

type Transformer interface {
	Start() error
	Transform(record Record) (Voice, error)
	Close() error
}
