package prot

type Speaker interface {
	Start() error
	Close() error
}
