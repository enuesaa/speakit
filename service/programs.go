package service

import (
	"github.com/enuesaa/speakit/repository"
)

type Program struct {
	Id string
	Title string
	Content string
}

type ProgramsService struct {
	repos repository.Repos
}

func NewProgramsService(repos repository.Repos) ProgramsService {
	return ProgramsService{
		repos,
	}
}

func (srv *ProgramsService) List() []Program {
	ids := srv.repos.Redis.Keys("programs:*")
	list := make([]Program, 0)

	for _, id := range ids {
		list = append(list, srv.Get(id))
	}
	return list
}

func (srv *ProgramsService) Get(id string) Program {
	value := srv.repos.Redis.Get("programs:" + id)
	return parseJson[Program](value)
}

func (srv *ProgramsService) Create(program Program) string {
	program.Id = createId()
	// srv.repos.Storage.Upload(id+".wav", body)
	srv.repos.Redis.Set("programs:" + program.Id, toJson(program))
	return program.Id
}

func (srv *ProgramsService) Download(id string) (string, error) {
	return srv.repos.Storage.Download(id + ".wav")
}
