package database

type ShortenRepository interface {
	// Add short url binding to long one
	AddUrl(shortUrl string, longUrl string) error
	// Check if short url exists
	Exists(shortUrl string) (bool, error)
	// Get returns long version of url
	Get(shortUrl string) (string, error)
}
