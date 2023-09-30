package service

import (
	"github.com/enuesaa/speakit/repository"
)

type Program struct {
	Id string
	Title string
	Content string
	Converted bool
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
		list = append(list, srv.Get(renameKeyToId(id)))
	}
	return list
}

// TODO: should return error
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

func (srv *ProgramService) Delete(id string) {
	srv.repos.Redis.Delete("programs:" + id)
}

func (srv *ProgramService) Upload(id string, body string) error {
	if err := srv.repos.Storage.Upload(id + ".wav", body); err != nil {
		return err
	}
	srv.AddConvertedFlag(id)
	return nil
}

func (srv *ProgramService) AddConvertedFlag(id string) {
	program := srv.Get(id)
	program.Converted = true
	srv.repos.Redis.Set("programs:" + id, toJson(program))
}

func (srv *ProgramService) Download(id string) (string, error) {
	return srv.repos.Storage.Download(id + ".wav")
}


func (srv *ProgramService) Convert(id string) (error) {
	program := srv.Get(id)

	audioquery, err := srv.repos.Voicevox.AudioQuery(program.Title)
	if err != nil {
		return err
	}
	converted, err := srv.repos.Voicevox.Synthesis(audioquery)
	if err != nil {
		return err
	}
	return srv.Upload(id, converted)
}
