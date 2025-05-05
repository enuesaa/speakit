package prot

type Callfn interface {
	StartUp() error
	Close() error
}
