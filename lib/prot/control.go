package prot

type Controller interface {
	StartUp(logger Logger) error
	Close() error
}
