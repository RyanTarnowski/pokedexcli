package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			clean_user_input := cleanInput(scanner.Text())

			if len(clean_user_input) > 0 {
				if command, ok := getCommands()[clean_user_input[0]]; ok {
					err := command.callback(&cfg)

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
