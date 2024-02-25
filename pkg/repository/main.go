package repository

import (
	"fmt"
)

type Env struct {
	REDIS_HOST   string
	VOICEVOX_HOST string
}

type Repos struct {
	Redis    RedisRepositoryInterface
	Storage  StorageRepositoryInterface
	Voicevox VoicevoxRepositoryInterface
	Fs       FsRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Redis: &RedisRepository{
			Addr: env.REDIS_HOST,
		},
		Storage: &StoragefsRepository{},
		Voicevox: &VoicevoxRepository{
			BaseUrl: fmt.Sprintf("https://%s", env.VOICEVOX_HOST),
		},
		Fs:       &FsRepository{},
	}
}
