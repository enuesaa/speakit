package prot

type Generator interface {
	StartUp(logger Logger) error
	Generate() (Record, error)
	Close() error
}
