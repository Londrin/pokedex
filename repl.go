package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/londrin/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
	pokedex          map[string]pokeapi.Pokemon
}

func start_repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command, See HELP for commands")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List location areas to navigate",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous maps unless at beginning on list",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore given <location> Example:'explore canalave-city-area'",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect a pokemon's stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemon in your Pokedex",
			callback:    commandPokedex,
		},
	}
}
