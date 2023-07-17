package redis

import (
	"context"
	"errors"
	"time"

	"github.com/deka-microservices/go-url-shortener/internal/consts"
	"github.com/deka-microservices/go-url-shortener/internal/database/dberrors"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisRepository struct {
	db *redis.ClusterClient
}

func NewRedisRepository() (*RedisRepository, error) {
	addresses := viper.GetStringSlice(consts.CONFIG_REDIS_ADDRESSES)

	if len(addresses) == 0 {
		return nil, ErrRedisAddressIsEmpty
	}

	options := redis.ClusterOptions{
		Addrs:    addresses,
		Password: viper.GetString(consts.CONFIG_REDIS_PASSWORD),
	}

	db := redis.NewClusterClient(&options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisRepository{
		db: db,
	}, nil
}

func (repo *RedisRepository) AddUrl(ctx context.Context, shortUrl string, longUrl string) error {
	err := repo.db.Set(ctx, shortUrl, longUrl, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
func (repo *RedisRepository) Exists(ctx context.Context, shortUrl string) (bool, error) {
	// TODO! Use redis cmd instead Get based

	_, err := repo.Get(ctx, shortUrl)
	if errors.Is(err, dberrors.ErrShortUrlNotFound) {
		return false, nil
	}

	return err == nil, err
}
func (repo *RedisRepository) Get(ctx context.Context, shortUrl string) (string, error) {
	status := repo.db.Get(ctx, shortUrl)

	if status.Err() == redis.Nil {
		return "", dberrors.ErrShortUrlNotFound
	} else if status.Err() != nil {
		return "", status.Err()
	}

	longUrl, _ := status.Result()

	return longUrl, nil

}
