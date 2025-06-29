package pokeapi

import (
	"net/http"
	"time"
)

// Client is the API client for PokeAPI.
type Client struct {
	HTTPClient *http.Client
	Cache      *Cache
}

// NewClient creates a new PokeAPI client.
func NewClient(httpTimeout, cacheInterval time.Duration) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: httpTimeout},
		Cache:      NewCache(cacheInterval),
	}
}
