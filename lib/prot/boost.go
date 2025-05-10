package prot

type Booster interface {
	Callfn
	Boost(original Record) bool
}
