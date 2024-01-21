package repository

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

func NewRepos(env Env) Repos {
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
