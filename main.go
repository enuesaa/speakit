package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fmt.Println("aaa")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://gigazine.net/news/rss_2.0/")
	fmt.Println(feed.Title)
}
