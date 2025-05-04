package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/klausks/go-pokedex/internal/pokecache"
)

const API_BASE_URL = "https://pokeapi.co/api/v2"

type PokeApiClient struct {
	cache      *pokecache.Cache
	httpClient *http.Client
}

func NewPokeApiClient() *PokeApiClient {
	cache := pokecache.NewCache(time.Minute)
	client := http.DefaultClient
	return &PokeApiClient{cache: cache, httpClient: client}
}

type LocationAreaResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (client *PokeApiClient) GetLocationAreaNames(pageUrl string) (names []string, previousPageUrl, nextPageUrl string, err error) {
	locationAreas, err := client.getLocationAreas(pageUrl)
	if err != nil {
		return nil, "", "", err
	}
	var locationAreaNames = make([]string, len(locationAreas.Results))
	for i, locationArea := range locationAreas.Results {
		locationAreaNames[i] = locationArea.Name
	}
	return locationAreaNames, locationAreas.Previous, locationAreas.Next, nil
}

func (client *PokeApiClient) getLocationAreas(pageUrl string) (LocationAreaResponse, error) {
	var url string
	if pageUrl != "" {
		url = pageUrl
	} else {
		endpoint := "location-area?offset=0&limit=20"
		url = fmt.Sprintf("%s/%s", API_BASE_URL, endpoint)
	}
	if cached, exists := client.cache.Get(url); exists {
		var cachedLocationAreaResponsePage LocationAreaResponse
		if err := json.Unmarshal(cached, &cachedLocationAreaResponsePage); err != nil {
			return LocationAreaResponse{}, err
		}
		return cachedLocationAreaResponsePage, nil
	}

	httpResBody, err := getLocationAreasFromApi(url)
	if err != nil {
		return LocationAreaResponse{}, nil
	}

	client.cache.Add(url, httpResBody)

	var locationAreas LocationAreaResponse
	json.Unmarshal(httpResBody, &locationAreas)
	return locationAreas, nil
}

func getLocationAreasFromApi(url string) ([]byte, error) {
	fmt.Println("Sending request:", url)
	httpRes, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	httpResBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	return httpResBody, nil
}
