package service

import (
	"github.com/enuesaa/speakit/repository"
)

type Program struct {
	Id string
	Title string
	Content string
}

type ProgramService struct {
	repos repository.Repos
}

func NewProgramService(repos repository.Repos) ProgramService {
	return ProgramService{
		repos,
	}
}

func (srv *ProgramService) List() []Program {
	ids := srv.repos.Redis.Keys("programs:*")
	list := make([]Program, 0)

	for _, id := range ids {
		list = append(list, srv.Get(id))
	}
	return list
}

func (srv *ProgramService) Get(id string) Program {
	value := srv.repos.Redis.Get("programs:" + id)
	return parseJson[Program](value)
}

func (srv *ProgramService) Create(program Program) string {
	program.Id = createId()
	// srv.repos.Storage.Upload(id+".wav", body)
	srv.repos.Redis.Set("programs:" + program.Id, toJson(program))
	return program.Id
}

func (srv *ProgramService) Download(id string) (string, error) {
	return srv.repos.Storage.Download(id + ".wav")
}
