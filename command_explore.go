package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if args[0] == "" {
		return fmt.Errorf("No area given: %s", args[0])
	}
	mons, err := cfg.pokeapiClient.ExploreLocationArea(args[0])
	if err != nil {
		fmt.Printf("[API] Unable to process: %s", err)
		return err
	}

	for _, name := range mons {
		fmt.Println(name)
	}
	return nil
}
