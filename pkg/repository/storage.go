package repository

type StorageRepositoryInterface interface {
	Upload(key string, value []byte) error
	Download(key string) ([]byte, error)
}

type StorageRepository struct {
	Items map[string][]byte
}

func (repo *StorageRepository) Upload(key string, value []byte) error {
	repo.Items[key] = value
	return nil
}

func (repo *StorageRepository) Download(key string) ([]byte, error) {
	return repo.Items[key], nil
}
