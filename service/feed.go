package service

import (
	"github.com/enuesaa/speakit/repository"
)

type Feed struct {
	Name string
	Url  string
}

type FeedService struct {
	repos repository.Repos
}

func NewFeedSevice(repos repository.Repos) FeedService {
	return FeedService{
		repos,
	}
}

func (srv *FeedService) List() []string {
	return srv.repos.Redis.Keys("feeds:*")
}

func (srv *FeedService) Get(id string) Feed {
	return *new(Feed)
}

func (srv *FeedService) Create(feed Feed) string {
	srv.repos.Redis.Set("feedss:"+feed.Name, feed.Url)
	return ""
}

func (srv *FeedService) Delete(id string) {}
