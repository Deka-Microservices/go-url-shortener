package redis

import "errors"

var (
	ErrRedisAddressIsEmpty = errors.New("redis address is not set")
)
