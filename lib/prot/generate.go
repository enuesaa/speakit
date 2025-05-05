package prot

type Generator interface {
	Callfn
	Generate() (Record, error)
}
