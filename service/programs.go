package service

import (
	"github.com/enuesaa/speakit/repository"
)

type ProgramsService struct {
	repos repository.Repos
}

func NewProgramsService(repos repository.Repos) ProgramsService {
	return ProgramsService{
		repos,
	}
}

func (srv *ProgramsService) List() []string {
	return srv.repos.Redis.Keys("programs:*")
}

func (srv *ProgramsService) Get(id string) Feed {
	return *new(Feed)
}

func (srv *ProgramsService) Create(key string, body string) {
	srv.repos.Minio.Upload(key, body)
}
