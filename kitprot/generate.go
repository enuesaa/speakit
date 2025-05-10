package kitprot

type Generator interface {
	Callfn
	Generate() (Record, error)
}
