package database

import (
	"errors"

	"github.com/deka-microservices/go-url-shortener/internal/database/inmemory"
	"github.com/deka-microservices/go-url-shortener/internal/database/redis"
	"github.com/rs/zerolog/log"
)

var DB ShortenRepository

func InitGlobalDatabase() error {
	redisRepo, err := redis.NewRedisRepository()
	if errors.Is(err, redis.ErrRedisAddressIsEmpty) {
		log.Warn().Err(err).Msg("switching to inmemory db")
		DB = inmemory.NewInMemoryRepo()
		return nil
	} else if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to redis")
	}

	DB = redisRepo
	return nil
}
