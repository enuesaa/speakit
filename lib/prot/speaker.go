package prot

import "time"

type Speaker interface {
	StartUp() error
	Speak(Record) (time.Duration, error)
	Stop() error
	Close() error
}
