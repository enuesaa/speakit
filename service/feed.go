package service

import (
	"github.com/enuesaa/speakit/repository"
)

type Feed struct {
	Id string
    Name string
	Url string
}

type FeedService struct {
	repos repository.Repos
}

func NewFeedSevice(repos repository.Repos) FeedService {
	return FeedService {
		repos,
	}
}

func (srv *FeedService) List() []Feed {
	list := make([]Feed, 0)
	return list
}

func (srv *FeedService) Get(id string) Feed {
	return *new(Feed)
}

func (srv *FeedService) Create(feed Feed) string {
	return ""
}

func (srv *FeedService) Delete(id string) {}
