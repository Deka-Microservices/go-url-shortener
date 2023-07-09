package database

import "github.com/deka-microservices/go-url-shortener/internal/database/inmemory"

var DB ShortenRepository

func InitGlobalDatabase() error {
	DB = inmemory.NewInMemoryRepo()
	return nil
}
