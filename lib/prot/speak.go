package prot

type Speaker interface {
	Start() error
	Next() error
	Stop() error
	Close() error
}

type SonosSpeaker struct {}

func (g *SonosSpeaker) Start() error {
	// listen
	return nil
}

func (g *SonosSpeaker) Next() error {
	return nil
}

func (g *SonosSpeaker) Stop() error {
	return nil
}

func (g *SonosSpeaker) Close() error {
	return nil
}
