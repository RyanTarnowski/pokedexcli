package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2/location-area"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(pageUrl *string, cache *Cache) (LocationAreas, error) {
	url := baseURL
	if pageUrl != nil {
		url = *pageUrl
	}

	location_areas := LocationAreas{}

	//check the cache
	val, ok := cache.Get(url)
	if ok {
		fmt.Println("+++ cache hit +++")

		err := json.Unmarshal(val, &location_areas)
		if err != nil {
			return location_areas, err
		}

		return location_areas, nil
	}

	fmt.Println("--- cache miss ---")
	res, err := http.Get(url)
	if err != nil {
		return location_areas, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return location_areas, err
	}
	if err != nil {
		return location_areas, err
	}

	err = json.Unmarshal(body, &location_areas)
	if err != nil {
		return location_areas, err
	}

	//add to cache
	cache.Add(url, body)
	return location_areas, nil
}
