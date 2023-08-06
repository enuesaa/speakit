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

// Perform like API Gateway
// URLのマッピングはGoのAppでいろいろ変えられるイメージ

// ## URL体系
// - GET /feeds ... feed list
// - GET /feeds/{id}
// - POST /feeds
// - DELETE /feeds/{id}

// - POST /jobs ... fetch rss feed and request to convert. 202 を返したい
// - GET /jobs
// - GET /jobs/{id}

// - GET /contents ... 一覧
// - GET /contents/{id} ... asset id を返す

// - GET /assets/{id}  ... wav file

// - GET /  ... admin
// - GET /feeds
// - GET /player
//   - start
//   - next
//   - prev
//   - stop