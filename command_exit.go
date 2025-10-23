package main

import (
	"fmt"
	"os"
)

func commandExit(config *LocationArea) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
