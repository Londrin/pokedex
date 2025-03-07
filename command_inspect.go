package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a pokemon name to inspect")
	}

	name := args[0]

	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %s\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, poke := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", poke.Stat.Name, poke.Base_stat)
	}
	fmt.Println("Types:")
	for _, poke := range pokemon.Types {
		fmt.Printf("  - %s\n", poke.Type.Name)
	}
	return nil
}
