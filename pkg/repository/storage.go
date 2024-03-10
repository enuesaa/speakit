package repository

type StorageRepositoryInterface interface {
	Upload(key string, value string) error
	Download(key string) (string, error)
}

type StorageRepository struct {
	Items map[string]string
}

func (repo *StorageRepository) Upload(key string, value string) error {
	repo.Items[key] = value
	return nil
}

func (repo *StorageRepository) Download(key string) (string, error) {
	return repo.Items[key], nil
}
