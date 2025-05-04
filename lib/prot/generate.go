package prot

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	ext "github.com/mmcdole/gofeed/extensions"
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
	fp := gofeed.NewParser()
	feeds, err := fp.ParseURL(g.Feed)
	if err != nil {
		return err
	}

	for i, item := range feeds.Items {
		if i > 1 {
			break
		}

		meta := g.flattenExtensions(item.Extensions)
		meta["link"] = item.Link
		meta["description"] = item.Description

		g.list = append(g.list, Record{
			Text: item.Title,
			Meta: meta,
		})
	}
	return nil
}

func (g *RSSFeedGenerator) flattenExtensions(extensions ext.Extensions) map[string]string {
	result := make(map[string]string)

	for namespace, values := range extensions {
		for property, value := range values {
			key := fmt.Sprintf("%s_%s", namespace, property)
			joined := ""
			for i, ext := range value {
				if i > 0 {
					joined += ","
				}
				joined += ext.Value
			}
			result[key] = joined
		}
	}
	return result
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
