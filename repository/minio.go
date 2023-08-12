package repository

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioRepositoryInterface interface {
	Bucket() string
	Upload(key string, obj string)
	Download(key string) (string, error)
}

type MinioRepository struct{}

func (repo *MinioRepository) Bucket() string {
	return os.Getenv("BUCKET")
}

func (repo *MinioRepository) Upload(key string, obj string) {
	ctx := context.Background()

	client, err := minio.New("minio:9000", &minio.Options{
		Creds: credentials.NewEnvMinio(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := strings.NewReader(obj)

	info, err := client.PutObject(ctx, repo.Bucket(), key, reader, reader.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}


func (repo *MinioRepository) Download(key string) (string, error) {
	ctx := context.Background()

	client, err := minio.New("minio:9000", &minio.Options{
		Creds: credentials.NewEnvMinio(),
	})
	if err != nil {
		return "", err
	}

	minioObj, _ := client.GetObject(ctx, repo.Bucket(), key, minio.GetObjectOptions{})
	obj, err := io.ReadAll(minioObj)
	if err != nil {
		return "", err
	}

	return string(obj), nil
}
