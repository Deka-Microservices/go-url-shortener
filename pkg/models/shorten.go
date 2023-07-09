package models

type ShortenRequest struct {
	Url string `json:"url" xml:"url" binding:"required"`
}

type ShortenResponse struct {
	Url string `json:"url" xml:"url"`
}
