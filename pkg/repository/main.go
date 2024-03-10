package repository

type Env struct {
	REDIS_HOST   string
	VOICEVOX_BASE_URL string
}

type Repos struct {
	Storage  StorageRepositoryInterface
	Voicevox VoicevoxRepositoryInterface
	Fs       FsRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Storage: &StoragefsRepository{},
		Voicevox: &VoicevoxRepository{
			BaseUrl: env.VOICEVOX_BASE_URL,
		},
		Fs: &FsRepository{},
	}
}
