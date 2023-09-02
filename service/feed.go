package service

import (
	"encoding/json"
	"github.com/enuesaa/speakit/repository"
	"github.com/google/uuid"
)

type Feed struct {
	Id string
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

func (srv *FeedService) Delete(id string) {}
