package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			clean_user_input := cleanInput(scanner.Text())

			if len(clean_user_input) > 0 {
				fmt.Printf("Your command was: %s\n", clean_user_input[0])
			}
		}
	}
}

func cleanInput(text string) []string {
	scrubbed_text := strings.ToLower(text)
	words := strings.Fields(scrubbed_text)

	return words
}
