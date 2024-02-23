package repository

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type StoragefsRepository struct {}

func (repo *StoragefsRepository) Upload(key string, value string) error {
	path := fmt.Sprintf("./tmp/%s", key)
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte(value)); err != nil {
		return err
	}
	return nil
}

func (repo *StoragefsRepository) Download(key string) (string, error) {
	path := fmt.Sprintf("./tmp/%s", key)
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	body, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
