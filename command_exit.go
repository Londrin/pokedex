package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	cfg.pokeapiClient.Cache.Stop()
	os.Exit(0)
	return nil
}
