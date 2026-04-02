package main

import (
	"errors"
	"fmt"
	"github.com/RyanTarnowski/pokedexcli/internal/pokeapi"
	"math/rand/v2"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	next          *string
	previous      *string
	cache         *pokeapi.Cache
	caughtPokemon map[string]pokeapi.PokemonInfo
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
		"explore": {
			name:        "explore",
			description: "Displays Pokemon of queried location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *config, args ...string) error {
	location_areas, err := pokeapi.GetLocationAreas(cfg.next, cfg.cache)
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	location_areas, err := pokeapi.GetLocationAreas(cfg.previous, cfg.cache)
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you have to enter a location name")
	}

	area_name := args[0]
	location_area_info, err := pokeapi.GetLocationAreaInfo(&area_name, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area_name)
	fmt.Println("Found Pokemon:")
	for _, pe := range location_area_info.PokemonEncounters {
		fmt.Printf(" - %s\n", pe.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you have to enter a pokemon name to catch it")
	}

	pokemon_name := args[0]
	pokemon_info, err := pokeapi.GetPokemonInfo(&pokemon_name, cfg.cache)
	if err != nil {
		return err
	}

	catch_threshold := 50
	catch_chance := rand.IntN(pokemon_info.BaseExperience)

	//fmt.Printf("Catch chance: %v\n", catch_chance)
	//fmt.Printf("BaseExperience %v\n", pokemon_info.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
	if catch_chance == pokemon_info.BaseExperience || catch_chance < catch_threshold {
		fmt.Printf("%s was caught!\n", pokemon_name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}
