package prot

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
)

type Generator interface {
	Generate() ([]Record, error)
}

type RSSFeedGenerator struct {
	Feed string
	Count int
}

func (g *RSSFeedGenerator) Generate() ([]Record, error) {
	feedsSrv := service.NewFeedSevice(repository.Repos{})
	feeds, err := feedsSrv.TryFetch(g.Feed)
	if err != nil {
		return nil, err
	}
	var list []Record

	for i, item := range feeds.Items {
		if i > g.Count {
			break
		}
		list = append(list, Record{
			Text: item.Title,
		})
	}
	return list, nil
}
