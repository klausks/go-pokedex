package cli

import (
	"fmt"

	"github.com/klausks/go-pokedex/internal/pokeapi"
)

type mapCommand struct {
	context       *ApiRequestContext
	pokeApiClient *pokeapi.PokeApiClient
}

func NewMapCommand(context *ApiRequestContext, client *pokeapi.PokeApiClient) mapCommand {
	return mapCommand{context: context, pokeApiClient: client}
}

func (c mapCommand) Name() string {
	return "map"
}

func (c mapCommand) Description() string {
	return "Shows the next page of location areas"
}

func (c mapCommand) Execute(args []string) error {
	locationAreaNames, previousPageUrl, nextPageUrl, err := c.pokeApiClient.GetLocationAreaNames(c.context.next)
	if err != nil {
		return err
	}
	c.context.update(previousPageUrl, nextPageUrl)
	for _, areaName := range locationAreaNames {
		fmt.Println(areaName)
	}
	return nil
}
