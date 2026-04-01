package main

import (
	"bufio"
	"fmt"
	"github.com/RyanTarnowski/pokedexcli/internal/pokeapi"
	"os"
	"strings"
	"time"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}
	const interval = 5 * time.Minute
	cache := pokeapi.NewCache(interval)
	cfg.cache = cache

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			clean_user_input := cleanInput(scanner.Text())

			if len(clean_user_input) > 0 {
				if command, ok := getCommands()[clean_user_input[0]]; ok {

					args := []string{}
					if len(clean_user_input) > 1 {
						args = clean_user_input[1:]
					}

					err := command.callback(&cfg, args...)

					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}

func cleanInput(text string) []string {
	scrubbed_text := strings.ToLower(text)
	words := strings.Fields(scrubbed_text)

	return words
}
