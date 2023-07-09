package inmemory

import (
	"errors"

	"github.com/deka-microservices/go-url-shortener/internal/database/dberrors"
)

type InMememoryRepo struct {
	mapping map[string]string
}

func NewInMemoryRepo() *InMememoryRepo {
	return &InMememoryRepo{
		mapping: make(map[string]string),
	}
}

func (repo *InMememoryRepo) AddUrl(shortUrl string, longUrl string) error {
	if exist, _ := repo.Exists(shortUrl); exist {
		return dberrors.ErrKeyHasAlreadyExists
	}

	repo.mapping[shortUrl] = longUrl
	return nil
}

func (repo *InMememoryRepo) Exists(shortUrl string) (bool, error) {
	_, err := repo.Get(shortUrl)
	if errors.Is(err, dberrors.ErrShortUrlNotFound) {
		return false, nil
	}

	return err == nil, err
}

func (repo *InMememoryRepo) Get(shortUrl string) (string, error) {
	shortUrl, ok := repo.mapping[shortUrl]
	if !ok {
		return "", dberrors.ErrShortUrlNotFound
	}

	return shortUrl, nil

}
