package dberrors

import "errors"

var (
	ErrShortUrlNotFound    = errors.New("short url not found")
	ErrKeyHasAlreadyExists = errors.New("trying to insert key that already exists")
)
