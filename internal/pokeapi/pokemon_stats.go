package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)

func (c *Client) PokemonStats(pokeName string) (ShallowPokemon, error) {
	if pokeName == "" {
		return ShallowPokemon{}, fmt.Errorf("There needs to be a Pokemon name")
	}
	
	url := baseURL + "/pokemon/" + pokeName

	bytes, exists := c.pokeCache.Get(url)
	pokemonResp := ShallowPokemon{}
	
	if exists {
		err := json.Unmarshal(bytes, &pokemonResp)
		if err != nil {
			return ShallowPokemon{}, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)
		
		if err != nil {
			return ShallowPokemon{}, err
		}
			
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ShallowPokemon{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return ShallowPokemon{}, err
		}

		err = json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return ShallowPokemon{}, err
		}
		c.pokeCache.Add(url, data)
	}
	return pokemonResp, nil
}
