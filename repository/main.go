package repository

import (
	"os"
)

type Repos struct {
	Redis    RedisRepositoryInterface
	Httpcall HttpcallRepositoryInterface
	Storage  StorageRepositoryInterface
}

func NewRealRepos() Repos {
	return Repos {
		Redis:    &RedisRepository{},
		Httpcall: &HttpcallRepository{},
		Storage: &StorageRepository{
			Bucket:   os.Getenv("MINIO_BUCKET"),
			Endpoint: os.Getenv("MINIO_ENDPOINT"),
		},
	}
}
