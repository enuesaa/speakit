package repository

type Env struct {
	MINIO_BUCKET string
	MINIO_HOST string
	REDIS_HOST string
	ADMIN_HOST string
}

type Repos struct {
	Redis    RedisRepositoryInterface
	Httpcall HttpcallRepositoryInterface
	Storage  StorageRepositoryInterface
}

func NewRepos(env Env) Repos {
	return Repos{
		Redis: &RedisRepository{
			Addr: env.REDIS_HOST,
		},
		Httpcall: &HttpcallRepository{},
		Storage: &StorageRepository{
			Bucket:   env.MINIO_BUCKET,
			Endpoint: env.MINIO_HOST,
		},
	}
}
