package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageUrl != nil {
		url = *pageUrl
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
	var locationResp RespShallowLocations
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locationResp, nil
}
