package repository

type Env struct {
	OPENAI_API_KEY string
}

type Repos struct {
	Data     DataRepositoryInterface
	Storage  StorageRepositoryInterface
	OpenAI   OpenAIRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Data: &DataRepository{
			Items: make(map[string]string),
		},
		Storage: &StorageRepository{
			Items: make(map[string][]byte),
		},
		OpenAI: &OpenAIRepository{
			APIKey: env.OPENAI_API_KEY,
		},
	}
}
