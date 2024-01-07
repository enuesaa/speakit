package repository

import (
	"context"
	"io"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageRepositoryInterface interface {
	Upload(key string, value string) error
	Download(key string) (string, error)
}

type StorageRepository struct {
	Bucket   string
	Endpoint string
}

func (repo *StorageRepository) client() (*minio.Client, error) {
	return minio.New(repo.Endpoint, &minio.Options{
		Creds: credentials.NewEnvMinio(),
	})
}

func (repo *StorageRepository) Upload(key string, value string) error {
	client, err := repo.client()
	if err != nil {
		return err
	}

	ctx := context.Background()
	reader := strings.NewReader(value)
	size := reader.Size()
	options := minio.PutObjectOptions{}
	if _, err := client.PutObject(ctx, repo.Bucket, key, reader, size, options); err != nil {
		return err
	}
	return nil
}

func (repo *StorageRepository) Download(key string) (string, error) {
	client, err := repo.client()
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	options := minio.GetObjectOptions{}
	obj, err := client.GetObject(ctx, repo.Bucket, key, options)
	if err != nil {
		return "", err
	}
	value, err := io.ReadAll(obj)
	if err != nil {
		return "", err
	}
	return string(value), nil
}
