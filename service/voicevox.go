package service

import (
	"fmt"
	"strings"

	"github.com/enuesaa/speakit/repository"
)

type VoicevoxService struct {
	repos repository.Repos
}

func NewVoicevoxService(repos repository.Repos) VoicevoxService {
	return VoicevoxService{
		repos,
	}
}

func (srv *VoicevoxService) AudioQuery(text string) {
	body, err := srv.repos.Httpcall.Post(
		"http://voicevox:50021/audio_query?speaker=1&text=" + text,
		strings.NewReader(""),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", body)
}
