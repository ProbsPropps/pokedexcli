package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, pokeName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	pokeResp, err := cfg.pokeapiClient.PokemonStats(pokeName)
	if err != nil {
		return fmt.Errorf("There was trouble looking up the Pokemon you typed in")
	}
	catchChance := rand.IntN(pokeResp.BaseExperience)
	if catchChance < 100 {
		fmt.Printf("%s was caught!\n", pokeName)
		cfg.caughtPokemon[pokeResp.Name] = pokeResp
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokeName)
	}

	return nil
}
