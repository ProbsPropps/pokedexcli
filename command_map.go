package main

import (
	"fmt"
	"errors"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}
	
	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		return errors.New("You are on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)
	if err != nil {
	return err
	}
	
	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
