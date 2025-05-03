package prot

type Controller interface {
	StartUp() error
	Prev() error
	Next() error
	Stop() error
	Close() error
}
