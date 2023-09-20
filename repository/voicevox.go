package repository

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type VoicevoxRepositoryInterface interface {
	Post(url string, body io.Reader) (string, error)
}

type VoicevoxRepository struct{}

func (repo *VoicevoxRepository) post(url string, body io.Reader) (string, error) {
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	return string(b), nil
}


func (repo *VoicevoxRepository) AudioQuery(text string) (string, error) {
	body, err := repo.post(
		"http://voicevox:50021/audio_query?speaker=2&text=" + url.QueryEscape(text),
		strings.NewReader(""),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}

func (repo *VoicevoxRepository) Synthesis(query string) (string, error) {
	body, err := repo.post(
		"http://voicevox:50021/synthesis?speaker=2&text=",
		strings.NewReader(query),
	)
	if err != nil {
		return "", fmt.Errorf("voicevox error %w", err)
	}
	return body, nil
}
