package main

import "fmt"

func commandInspect(cfg *config, pokeName string) error {
	_, exists := cfg.caughtPokemon[pokeName]
	if !exists {
		fmt.Println("You have not caught that Pokemon")
	}
	pokemon := cfg.caughtPokemon[pokeName]

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	
	stats := pokemon.Stats

	fmt.Println("Stats:")
	for _, stat := range stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	
	if len(pokemon.Types) == 1 {
		fmt.Println("Type:")
	}else { fmt.Println("Types:")}
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
