package pokeapi

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Count 			int `json:"count"`
	NextURL 		*string `json:"next"`
	PreviousURL 	*string `json:"previous"`
	Results		 	[]LocationData `json:"results"`
}

type LocationData struct {
	Name 	string `json:"name"`
	URL 	string `json:"url"`
}


func GetLocations(url string) (LocationArea, error){
	var locations LocationArea
	res, err := http.Get(url)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return locations, fmt.Errorf("current status code is %d", res.StatusCode)
	}
	if err != nil {
		return locations, err
	}

	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, err
	}

	return locations, nil

}	
