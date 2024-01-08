package service

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
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

func (srv *FeedService) List() []Feed {
	keys := srv.repos.Redis.Keys("feeds:*")
	list := make([]Feed, 0)

	for _, key := range keys {
		list = append(list, srv.Get(renameKeyToId(key)))
	}
	return list
}

func (srv *FeedService) Get(id string) Feed {
	value := srv.repos.Redis.Get("feeds:" + id)

	return parseJson[Feed](value)
}

func (srv *FeedService) Create(feed Feed) string {
	feed.Id = createId()
	srv.repos.Redis.Set("feeds:"+feed.Id, toJson(feed))
	return feed.Id
}

func (srv *FeedService) Delete(id string) {
	srv.repos.Redis.Delete("feeds:" + id)
}

func (srv *FeedService) Refetch(id string) (*Realfeed, error) {
	feed := srv.Get(id)
	url := feed.Url

	fp := gofeed.NewParser()
	realfeed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return &Realfeed{*realfeed}, nil
}

func (srv *FeedService) TryFetch(url string) (Realfeed, error) {
	fp := gofeed.NewParser()
	realfeed, err := fp.ParseURL(url)
	if err != nil {
		return Realfeed{}, err
	}

	return Realfeed{*realfeed}, nil
}
