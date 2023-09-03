package service

import (
	"encoding/json"

	"github.com/enuesaa/speakit/repository"
	"github.com/google/uuid"
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

func (srv *ProgramsService) ListKeys() []string {
	return srv.repos.Redis.Keys("programs:*")
}

func (srv *ProgramsService) List() []Program {
	ids := srv.ListKeys()
	list := make([]Program, 0)

	for _, id := range ids {
		list = append(list, srv.Get(id))
	}
	return list
}

func (srv *ProgramsService) Get(id string) Program {
	return *new(Program)
}

func (srv *ProgramsService) Create(program Program) string {
	uid, _ := uuid.NewUUID()
	id := uid.String()
	program.Id = id
	bfeed, _ := json.Marshal(program)

	// srv.repos.Storage.Upload(id+".wav", body)
	srv.repos.Redis.Set("programs:" + id, string(bfeed))
	return id
}

func (srv *ProgramsService) Download(id string) (string, error) {
	return srv.repos.Storage.Download(id + ".wav")
}
