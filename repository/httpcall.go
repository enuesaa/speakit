package repository

import (
	"io"
	"net/http"
)

// Deprecated: instead, please use voicevox repository
type HttpcallRepositoryInterface interface {
	Post(url string, body io.Reader) (string, error)
}

type HttpcallRepository struct{}

func (repo *HttpcallRepository) Post(url string, body io.Reader) (string, error) {
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)

	return string(b), nil
}
