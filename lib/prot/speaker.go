package prot

type Speaker interface {
	Start() error
	Next(Record) error
	Stop() error
	Close() error
}
