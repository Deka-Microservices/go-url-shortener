package database

import "context"

type ShortenRepository interface {
	// Add short url binding to long one
	AddUrl(ctx context.Context, string, longUrl string) error
	// Check if short url exists
	Exists(ctx context.Context, shortUrl string) (bool, error)
	// Get returns long version of url
	Get(ctx context.Context, shortUrl string) (string, error)
}
