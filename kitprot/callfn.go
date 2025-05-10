package kitprot

type Callfn interface {
	StartUp() error
	Close() error
}
