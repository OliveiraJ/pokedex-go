package pokeapi

import (
	"io"
	"net/http"
)

type Cache interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
}

type Client struct {
	cache      Cache
	httpClient *http.Client
}

func NewClient(cache Cache) *Client {
	return &Client{
		cache:      cache,
		httpClient: http.DefaultClient,
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	if body, ok := c.cache.Get(url); ok {
		return body, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)
	return body, nil
}
