package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

func main() {
	fmt.Println("aaa")

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://gigazine.net/news/rss_2.0/")
	fmt.Println(feed.Title)

	app := fiber.New()
	// app.Get("/api", handler.ListContents)
	app.Static("/", "./public")
	app.Listen(":3000")
}
