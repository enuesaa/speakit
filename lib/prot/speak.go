package prot

type Speaker interface {
	Start() error
	Close() error
}

type SonosSpeaker struct {}

func (g *SonosSpeaker) Start() error {
	return nil
}

func (g *SonosSpeaker) Close() error {
	return nil
}
