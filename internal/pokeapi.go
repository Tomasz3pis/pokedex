package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Tomasz3pis/pokedex/internal/pokedex"
)

type RespShallowPokemons struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"encounter_method,omitempty"`
		VersionDetails []struct {
			Rate    int `json:"rate,omitempty"`
			Version struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"encounter_method_rates,omitempty"`
	GameIndex int `json:"game_index,omitempty"`
	ID        int `json:"id,omitempty"`
	Location  struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"location,omitempty"`
	Name  string `json:"name,omitempty"`
	Names []struct {
		Language struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"language,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"names,omitempty"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name,omitempty"`
			URL  string `json:"url,omitempty"`
		} `json:"pokemon,omitempty"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance,omitempty"`
				ConditionValues []any `json:"condition_values,omitempty"`
				MaxLevel        int   `json:"max_level,omitempty"`
				Method          struct {
					Name string `json:"name,omitempty"`
					URL  string `json:"url,omitempty"`
				} `json:"method,omitempty"`
				MinLevel int `json:"min_level,omitempty"`
			} `json:"encounter_details,omitempty"`
			MaxChance int `json:"max_chance,omitempty"`
			Version   struct {
				Name string `json:"name,omitempty"`
				URL  string `json:"url,omitempty"`
			} `json:"version,omitempty"`
		} `json:"version_details,omitempty"`
	} `json:"pokemon_encounters,omitempty"`
}

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetPokemon(name string) (pokedex.Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	var pokemon pokedex.Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	return pokemon, nil
}

func (c *Client) ListPokemons(loc string) (RespShallowPokemons, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + loc
	if data, exists := c.cache.Get(loc); exists {
		var pokemonResp RespShallowPokemons
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return RespShallowPokemons{}, err
		}
		return pokemonResp, nil

	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	var pokemonResp RespShallowPokemons
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespShallowPokemons{}, err
	}
	c.cache.Add(loc, data)
	return pokemonResp, nil
}

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	var locationResp RespShallowLocations

	if data, exists := c.cache.Get(url); exists {
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, data)
	return locationResp, nil
}
