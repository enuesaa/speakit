package repository

type Env struct {
	VOICEVOX_BASE_URL string
}

type Repos struct {
	Data     DataRepositoryInterface
	Storage  StorageRepositoryInterface
	Voicevox VoicevoxRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Data: &DataRepository{},
		Storage: &StorageRepository{},
		Voicevox: &VoicevoxRepository{
			BaseUrl: env.VOICEVOX_BASE_URL,
		},
	}
}
