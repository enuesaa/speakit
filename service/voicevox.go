package service

import (
	"fmt"
	"net/url"
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

func (srv *VoicevoxService) AudioQuery(text string) (string, error) {
	body, err := srv.repos.Httpcall.Post(
		"http://voicevox:50021/audio_query?speaker=1&text=" + url.QueryEscape(text),
		strings.NewReader(""),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}

func (srv *VoicevoxService) Synthesis(query string) (string, error) {
	body, err := srv.repos.Httpcall.Post(
		"http://voicevox:50021/synthesis?speaker=1&text=",
		strings.NewReader(query),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}
