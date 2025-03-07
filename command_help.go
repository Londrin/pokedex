package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, comm := range getCommands() {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}
