package kitprot

type Transformer interface {
	Callfn
	Transform(record *Record) error
}
