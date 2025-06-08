package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)


func (c *Client) ExploreLocation(locationName *string) (ShallowExplore, error) {
	if locationName == nil {
		return ShallowExplore{}, fmt.Errorf("There needs to be a location name")
	}

	url := baseURL + "/location-area/" + *locationName

	bytes, exists := c.pokeCache.Get(url)
	exploreResp := ShallowExplore{}

	if exists {
		err := json.Unmarshal(bytes, &exploreResp)
		if err != nil {
			return ShallowExplore{}, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ShallowExplore{}, err
		}
	
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ShallowExplore{}, err
		}
		defer resp.Body.Close()
	
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return ShallowExplore{}, err
		}
	
		err = json.Unmarshal(data, &exploreResp)
		if err != nil {
			return ShallowExplore{}, err
		}
		c.pokeCache.Add(url, data)
	}
	return exploreResp, nil
}
