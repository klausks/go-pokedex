package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API_BASE_URL = "https://pokeapi.co/api/v2/"

type LocationAreaBatch struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaNames(pageUrl string) (LocationAreaBatch, error) {
	var url string
	if pageUrl != "" {
		url = pageUrl
	} else {
		endpoint := "location-area"
		url = fmt.Sprintf("%s/%s", API_BASE_URL, endpoint)
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaBatch{}, err
	}
	var locationAreas LocationAreaBatch
	json.NewDecoder(res.Body).Decode(&locationAreas)

	return locationAreas, nil
}
