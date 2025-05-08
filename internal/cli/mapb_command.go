package cli

import (
	"fmt"

	"github.com/klausks/go-pokedex/internal/pokeapi"
)

type mapbCommand struct {
	context       *ApiRequestContext
	pokeApiClient *pokeapi.PokeApiClient
}

func NewMapbCommand(context *ApiRequestContext, apiClient *pokeapi.PokeApiClient) mapbCommand {
	return mapbCommand{context: context, pokeApiClient: apiClient}
}

func (c mapbCommand) Name() string {
	return "mapb"
}

func (c mapbCommand) Description() string {
	return "Shows the previous page of location areas"
}

func (c mapbCommand) Execute(args []string) error {
	locationAreaNames, previousPageUrl, nextPageUrl, err := c.pokeApiClient.GetLocationAreaNames(c.context.previous)
	if err != nil {
		return err
	}
	c.context.update(previousPageUrl, nextPageUrl)
	for _, areaName := range locationAreaNames {
		fmt.Println(areaName)
	}
	return nil
}
