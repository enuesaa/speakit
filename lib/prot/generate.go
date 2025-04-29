package prot

type Generator interface {
	Start() error
	Generate() (Record, error)
	Close() error
}
