package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
	}
}

func cleanInput(text string) []string {
	strLowered := strings.ToLower(text)
	words := strings.Fields(strLowered)
	return words
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}
