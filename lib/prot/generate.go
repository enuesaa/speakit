package prot

type Generator interface {
	StartUp() error
	Generate() (Record, error)
	Close() error
}
