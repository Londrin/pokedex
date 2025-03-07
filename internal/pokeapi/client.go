package pokeapi

import (
	"time"

	"github.com/londrin/pokedex/internal/pokecache"
)

type Client struct {
	Cache   pokecache.Cache
	pokedex map[string]Pokemon
}

func NewClient(interval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(interval),
	}
}
