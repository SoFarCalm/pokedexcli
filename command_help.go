package main

import "fmt"

func commandHelp(config *LocationArea) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
