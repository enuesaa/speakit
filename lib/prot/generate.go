package prot

type Generator interface {
	Generate() ([]Record, error)
}

type RSSFeedGenerator struct {}

func (g *RSSFeedGenerator) Generate() ([]Record, error) {
	return nil, nil
}
