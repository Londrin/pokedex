package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("go catch pokemon")
	}

	fmt.Println("Your Pokedex:")
	for poke := range cfg.pokedex {
		fmt.Printf("  - %s\n", poke)
	}

	return nil
}
