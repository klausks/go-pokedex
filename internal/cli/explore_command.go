package cli

import (
	"fmt"

	"github.com/klausks/go-pokedex/internal/pokeapi"
)

type exploreCommand struct {
	pokeApiClient *pokeapi.PokeApiClient
}

func NewExploreCommand(client *pokeapi.PokeApiClient) exploreCommand {
	return exploreCommand{pokeApiClient: client}
}

func (c exploreCommand) Name() string {
	return "explore"
}

func (c exploreCommand) Description() string {
	return "Shows details about a location area"
}

func (c exploreCommand) Execute(args []string) error {
	if args == nil || len(args) != 1 {
		return fmt.Errorf("correct usage: 'explore <area_name>'")
	}
	locationAreaName := args[0]
	pokemonNames, err := c.pokeApiClient.GetLocationAreaPokemonEncounters(locationAreaName)
	if err != nil {
		return err
	}
	for _, areaName := range pokemonNames {
		fmt.Println(areaName)
	}
	return nil
}
