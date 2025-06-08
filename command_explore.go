package main

import "fmt"

func commandExplore(cfg *config, locationName string) error {
	fmt.Printf("Exploring %s...\n", locationName)
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(&locationName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon!")
	for _, pok := range exploreResp.PokemonEncounters {
		fmt.Printf("- %s\n", pok.Pokemon.Name)
	}
	return nil
}
