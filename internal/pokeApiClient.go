package pokeClient

import (
    "net/http"
    "fmt"
)

type Config struct {
	Next     string `json:"next,omitempty"`
	Previous any    `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
	} `json:"results,omitempty"`
}

func NewConfig() Config, Error {
    url := "https://pokeapi.co/api/v2/location-area"
    res, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("Failed to get poke config: %w", err)
    }
    defer res.Body.Close()
    var cfg Config
    err = json.Unmarshal(res.Body, &cfg)
    if err != nil {
        return nil, fmt.Errorf("Failed to unmarshal response: %w", err)
    }
    return cfg, nil    
}
