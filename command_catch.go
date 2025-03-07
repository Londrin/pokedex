package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetCatchPokemon(name)
	if err != nil {
		return err
	}

	result := rand.Intn(pokemon.Base_experience)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if result > 50 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
