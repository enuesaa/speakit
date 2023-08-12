package service

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/google/uuid"
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

func (srv *ProgramsService) Create(body string) {
	uid, _ := uuid.NewUUID()
	id := uid.String()

	srv.repos.Minio.Upload(id + ".wav", body)
	srv.repos.Redis.Set("programs:" + id, "")
}

func (srv *ProgramsService) Download(id string) (string, error) {
	return srv.repos.Minio.Download(id + ".wav")
}
