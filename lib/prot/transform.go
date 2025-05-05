package prot

type Transformer interface {
	Callfn
	Transform(record *Record) error
}
