package model

type AddShortUrlResponse struct {
	OrgUrl   string `json:"originalUrl"`
	ShortUrl string `json:"shortUrl"`
}
