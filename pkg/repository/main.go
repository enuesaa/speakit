package repository

import "os"

type Env struct {
	MINIO_BUCKET string
	MINIO_HOST   string
	REDIS_HOST   string
	VOICEVOX_BASE_URL string
}

type Repos struct {
	Redis    RedisRepositoryInterface
	Storage  StorageRepositoryInterface
	Voicevox VoicevoxRepositoryInterface
	Fs       FsRepositoryInterface
}

func NewRepos() Repos {
	env := Env{
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST:   os.Getenv("MINIO_HOST"),
		REDIS_HOST:   os.Getenv("REDIS_HOST"),
		VOICEVOX_BASE_URL: os.Getenv("VOICEVOX_BASE_URL"),
	}

	return Repos{
		Redis: &RedisRepository{
			Addr: env.REDIS_HOST,
		},
		Storage: &StorageRepository{
			Bucket:   env.MINIO_BUCKET,
			Endpoint: env.MINIO_HOST,
		},
		Voicevox: &VoicevoxRepository{
			BaseUrl: env.VOICEVOX_BASE_URL,
		},
		Fs:       &FsRepository{},
	}
}
