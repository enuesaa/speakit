package repository

import (
	"strings"
)

type DataRepositoryInterface interface {
	Keys(pattern string) []string
	Get(key string) string
	Set(key string, value string)
	Delete(key string)
}

type DataRepository struct {
	Items map[string]string
}

func (repo *DataRepository) Keys(pattern string) []string {
	list := make([]string, 0)
	for key, _ := range repo.Items {
		if strings.HasPrefix(key, pattern) {
			list = append(list, key)
		}
	}
	return list
}

func (repo *DataRepository) Get(key string) string {
	return repo.Items[key]
}

func (repo *DataRepository) Set(key string, value string) {
	repo.Items[key] = value
}

func (repo *DataRepository) Delete(key string) {
	delete(repo.Items, key)
}
