package service

import (
	"fmt"
	"io"

	"github.com/enuesaa/speakit/pkg/repository"
)

type Program struct {
	Id        string
	Title     string
	Content   string
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
	ids := srv.repos.Data.Keys("programs:")
	list := make([]Program, 0)

	for _, id := range ids {
		list = append(list, srv.Get(renameKeyToId(id)))
	}
	return list
}

// TODO: should return error
func (srv *ProgramService) Get(id string) Program {
	value := srv.repos.Data.Get("programs:" + id)
	return parseJson[Program](value)
}

func (srv *ProgramService) Create(program Program) string {
	program.Id = createId()
	// srv.repos.Storage.Upload(id+".mp3", body)
	srv.repos.Data.Set("programs:"+program.Id, toJson(program))
	return program.Id
}

func (srv *ProgramService) Delete(id string) {
	srv.repos.Data.Delete("programs:" + id)
}

func (srv *ProgramService) Upload(id string, body []byte) error {
	if err := srv.repos.Storage.Upload(id, body); err != nil {
		return err
	}
	srv.AddConvertedFlag(id)
	return nil
}

func (srv *ProgramService) AddConvertedFlag(id string) {
	program := srv.Get(id)
	program.Converted = true
	fmt.Printf("flag: %+v\n", program)
	srv.repos.Data.Set("programs:"+id, toJson(program))
}

func (srv *ProgramService) Download(id string) ([]byte, error) {
	return srv.repos.Storage.Download(id)
}

func (srv *ProgramService) Convert(id string) error {
	program := srv.Get(id)

	converted, err := srv.repos.OpenAI.Speech(program.Title)
	if err != nil {
		return err
	}
	fmt.Printf("converted: %s\n", id)

	data, err := io.ReadAll(converted)
	if err != nil {
		return err
	}

	return srv.Upload(id, data)
}
