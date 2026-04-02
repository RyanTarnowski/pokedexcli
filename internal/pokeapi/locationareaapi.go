package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	locationareabaseURL = "https://pokeapi.co/api/v2/location-area/"
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

type LocationAreaInfo struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(pageUrl *string, cache *Cache) (LocationAreas, error) {
	url := locationareabaseURL
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
			return LocationAreas{}, err
		}

		return location_areas, nil
	}

	fmt.Println("--- cache miss ---")
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreas{}, err
	}
	if err != nil {
		return LocationAreas{}, err
	}

	err = json.Unmarshal(body, &location_areas)
	if err != nil {
		return LocationAreas{}, err
	}

	//add to cache
	cache.Add(url, body)
	return location_areas, nil
}

func GetLocationAreaInfo(LocationArea *string, cache *Cache) (LocationAreaInfo, error) {
	url := locationareabaseURL + *LocationArea

	location_area_info := LocationAreaInfo{}

	//check the cache
	val, ok := cache.Get(url)
	if ok {
		fmt.Println("+++ cache hit +++")

		err := json.Unmarshal(val, &location_area_info)
		if err != nil {
			return LocationAreaInfo{}, err
		}

		return location_area_info, nil
	}

	fmt.Println("--- cache miss ---")
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreaInfo{}, err
	}
	if err != nil {
		return LocationAreaInfo{}, err
	}

	err = json.Unmarshal(body, &location_area_info)
	if err != nil {
		return LocationAreaInfo{}, err
	}

	//add to cache
	cache.Add(url, body)
	return location_area_info, nil
}
