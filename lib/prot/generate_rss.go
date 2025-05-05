package prot

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
	ext "github.com/mmcdole/gofeed/extensions"
)

type RSSGenerator struct {
	Feed string

	log  LogBehavior
	list []Record
}

func (g *RSSGenerator) Inject(log LogBehavior) {
	g.log = log
}

func (g *RSSGenerator) StartUp() error {
	fp := gofeed.NewParser()
	feeds, err := fp.ParseURL(g.Feed)
	if err != nil {
		return err
	}

	for _, item := range feeds.Items {
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

func (g *RSSGenerator) flattenExtensions(extensions ext.Extensions) map[string]string {
	result := make(map[string]string)

	for namespace, values := range extensions {
		for property, value := range values {
			key := fmt.Sprintf("%s%s%s", namespace, strings.ToUpper(property[:1]), strings.ToLower(property[1:]))
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

func (g *RSSGenerator) Generate() (Record, error) {
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

func (g *RSSGenerator) Close() error {
	return nil
}
