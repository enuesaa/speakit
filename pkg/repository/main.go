package repository

import "os"

type Env struct {
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
		REDIS_HOST:   os.Getenv("REDIS_HOST"),
		VOICEVOX_BASE_URL: os.Getenv("VOICEVOX_BASE_URL"),
	}

	return Repos{
		Redis: &RedisRepository{
			Addr: env.REDIS_HOST,
		},
		Storage: &StoragefsRepository{},
		Voicevox: &VoicevoxRepository{
			BaseUrl: env.VOICEVOX_BASE_URL,
		},
		Fs:       &FsRepository{},
	}
}
