package repository

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type VoicevoxRepositoryInterface interface {
	SetBaseUrl(url string)
	Post(url string, body io.Reader) (string, error)
	AudioQuery(text string) (string, error)
	Synthesis(query string) (string, error)
}

type VoicevoxRepository struct {
	BaseUrl string
}

func (repo *VoicevoxRepository) SetBaseUrl(url string) {
	repo.BaseUrl = url
}

func (repo *VoicevoxRepository) Post(url string, body io.Reader) (string, error) {
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	return string(b), nil
}

func (repo *VoicevoxRepository) AudioQuery(text string) (string, error) {
	body, err := repo.Post(
		repo.BaseUrl+"/audio_query?speaker=2&text="+url.QueryEscape(text),
		strings.NewReader(""),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}

func (repo *VoicevoxRepository) Synthesis(query string) (string, error) {
	body, err := repo.Post(
		repo.BaseUrl+"/synthesis?speaker=2&text=",
		strings.NewReader(query),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}
