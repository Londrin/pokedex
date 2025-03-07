package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocation)
	if err != nil {
		return fmt.Errorf("unable to find map: %s", err)
	}

	cfg.nextLocation = locations.Next
	cfg.previousLocation = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocation == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.previousLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation = locations.Next
	cfg.previousLocation = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
