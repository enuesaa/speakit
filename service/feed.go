package service

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/speakit/repository"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Id string
	Name string
	Url  string
}
type Realfeed struct {
	gofeed.Feed
}

type FeedService struct {
	repos repository.Repos
}

func NewFeedSevice(repos repository.Repos) FeedService {
	return FeedService{
		repos,
	}
}

func (srv *FeedService) ListKeys() []string {
	return srv.repos.Redis.Keys("feeds:*")
}

func (srv *FeedService) List() []Feed {
	ids := srv.ListKeys()
	list := make([]Feed, 0)

	for _, id := range ids {
		list = append(list, srv.Get(id))
	}
	return list
}

func (srv *FeedService) Get(id string) Feed {
	value := srv.repos.Redis.Get("feeds:" + id)
	fmt.Println(value)

	return *new(Feed)
}

func (srv *FeedService) Create(feed Feed) string {
	uid, _ := uuid.NewUUID()
	id := uid.String()
	feed.Id = id
	bfeed, _ := json.Marshal(feed)
	srv.repos.Redis.Set("feeds:" + id, string(bfeed))
	return id
}

func (srv *FeedService) Delete(id string) {
	srv.repos.Redis.Delete("feeds:" + id)
}

func (srv *FeedService) Refetch(id string) Realfeed {
	feed := srv.Get(id)
	url := feed.Url

	fp := gofeed.NewParser()
	realfeed, _ := fp.ParseURL(url)

	return Realfeed{ *realfeed }
}
