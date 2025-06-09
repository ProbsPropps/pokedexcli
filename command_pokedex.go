package main

import "fmt"

func commandPokedex(cfg *config, na string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Go catch some Pokemon!")
		return nil
	}
	
	fmt.Println("Your Pokedex:")

	for pokemon := range cfg.caughtPokemon {
		fmt.Printf("  -%s\n", pokemon)
	}

	return nil

}
