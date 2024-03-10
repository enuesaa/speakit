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
		Data: &DataRepository{
			Items: make(map[string]string),
		},
		Storage: &StorageRepository{
			Items: make(map[string]string),
		},
		Voicevox: &VoicevoxRepository{
			BaseUrl: env.VOICEVOX_BASE_URL,
		},
	}
}
