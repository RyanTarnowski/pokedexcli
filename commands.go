package main

import (
	"fmt"
	"github.com/RyanTarnowski/pokedexcli/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     *string
	previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous location areas",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	location_areas, err := pokeapi.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = location_areas.Next
	cfg.previous = location_areas.Previous

	for _, la := range location_areas.Results {
		fmt.Println(la.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	location_areas, err := pokeapi.GetLocationAreas(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = location_areas.Next
	cfg.previous = location_areas.Previous

	for _, la := range location_areas.Results {
		fmt.Println(la.Name)
	}

	return nil
}
