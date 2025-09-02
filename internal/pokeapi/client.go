package pokeapi

import (
	"net/http"
	"time"

	"github.com/ty-strong/pokedexcli/internal/cache"
)

// Client -
type Client struct {
	cache      *cache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
