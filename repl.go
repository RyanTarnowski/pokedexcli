package main

import (
	"strings"
)

func cleanInput(text string) []string {
	scrubbed_text := strings.ToLower(text)
	words := strings.Fields(scrubbed_text)

	return words
}
