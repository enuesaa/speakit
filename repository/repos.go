package repository

type Repos struct {
	Redis    RedisRepositoryInterface
	Httpcall HttpcallRepositoryInterface
	Storage  StorageRepositoryInterface
}