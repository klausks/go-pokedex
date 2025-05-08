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

type LocationAreasResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaPokemonEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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

func (client *PokeApiClient) GetLocationAreaPokemonEncounters(locationAreaName string) (pokemonNames []string, err error) {
	locationAreaEncounters, err := client.getLocationAreaEncounters(locationAreaName)
	if err != nil {
		return nil, err
	}
	var encounteredPokemonNames = make([]string, len(locationAreaEncounters.PokemonEncounters))
	for i, pokemonEncounter := range locationAreaEncounters.PokemonEncounters {
		encounteredPokemonNames[i] = pokemonEncounter.Pokemon.Name
	}
	return encounteredPokemonNames, nil
}

func (client *PokeApiClient) getLocationAreas(pageUrl string) (LocationAreasResponse, error) {
	var url string
	if pageUrl != "" {
		url = pageUrl
	} else {
		endpoint := "location-area?offset=0&limit=20"
		url = fmt.Sprintf("%s/%s", API_BASE_URL, endpoint)
	}
	if cached, exists := client.cache.Get(url); exists {
		var cachedLocationAreaResponsePage LocationAreasResponse
		if err := json.Unmarshal(cached, &cachedLocationAreaResponsePage); err != nil {
			return LocationAreasResponse{}, err
		}
		return cachedLocationAreaResponsePage, nil
	}

	httpResBody, err := callGet(url)
	if err != nil {
		return LocationAreasResponse{}, nil
	}

	client.cache.Add(url, httpResBody)

	var locationAreas LocationAreasResponse
	json.Unmarshal(httpResBody, &locationAreas)
	return locationAreas, nil
}

func (client *PokeApiClient) getLocationAreaEncounters(locationAreaName string) (LocationAreaPokemonEncounters, error) {
	var url string
	endpoint := "location-area"
	queryString := "?offset=0&limit=20"
	url = fmt.Sprintf("%s/%s/%s%s", API_BASE_URL, endpoint, locationAreaName, queryString)

	if cached, exists := client.cache.Get(url); exists {
		var cachedEncounters LocationAreaPokemonEncounters
		if err := json.Unmarshal(cached, &cachedEncounters); err != nil {
			return LocationAreaPokemonEncounters{}, err
		}
		return cachedEncounters, nil
	}

	httpResBody, err := callGet(url)
	if err != nil {
		return LocationAreaPokemonEncounters{}, err
	}

	client.cache.Add(url, httpResBody)

	var locationAreaPokemonEncounters LocationAreaPokemonEncounters
	if err := json.Unmarshal(httpResBody, &locationAreaPokemonEncounters); err != nil {
		return LocationAreaPokemonEncounters{}, err
	}
	return locationAreaPokemonEncounters, nil
}

func callGet(url string) ([]byte, error) {
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

	if httpRes.StatusCode > 299 {
		return nil, fmt.Errorf("received non-success responde code %d, response body: %s", httpRes.StatusCode, httpResBody)
	}

	return httpResBody, nil
}
