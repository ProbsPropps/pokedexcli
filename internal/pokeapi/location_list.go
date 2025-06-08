package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	bytes, exists := c.pokeCache.Get(url)
	locationsResp := ShallowLocations{}

	if exists {
		err := json.Unmarshal(bytes, &locationsResp)
		if err != nil {
			return ShallowLocations{}, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ShallowLocations{}, err
		}
	
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ShallowLocations{}, err
		}
		defer resp.Body.Close()
	
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return ShallowLocations{}, err
		}
	
		err = json.Unmarshal(data, &locationsResp)
		if err != nil {
			return ShallowLocations{}, err
		}
		c.pokeCache.Add(url, data)
	}
	return locationsResp, nil
}


	
