package prot

import (
	"fmt"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
)

type Generator interface {
	StartUp() error
	Generate() (Record, error)
	Close() error
}

type RSSFeedGenerator struct {
	Feed string

	list []Record
}

func (g *RSSFeedGenerator) StartUp() error {
	feedsSrv := service.NewFeedSevice(repository.Repos{})
	feeds, err := feedsSrv.TryFetch(g.Feed)
	if err != nil {
		return err
	}

	for i, item := range feeds.Items {
		if i > 1 {
			break
		}
		g.list = append(g.list, Record{
			Text: item.Title,
		})
	}
	return nil
}

func (g *RSSFeedGenerator) Generate() (Record, error) {
	if len(g.list) > 0 {
		first := g.list[0]
		if len(g.list) > 1 {
			g.list = g.list[1:]
		} else {
			g.list = []Record{}
		}
		return first, nil
	}
	return Record{}, fmt.Errorf("final")
}

func (g *RSSFeedGenerator) Close() error {
	return nil
}
