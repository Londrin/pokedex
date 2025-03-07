package main

import (
	"time"

	"github.com/londrin/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(1 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}
	start_repl(cfg)
}
