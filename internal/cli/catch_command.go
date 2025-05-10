package cli

import (
	"fmt"
	"math/rand/v2"

	"github.com/klausks/go-pokedex/internal/pokeapi"
	"github.com/klausks/go-pokedex/model"
)

type catchCommand struct {
	pokeApiClient  *pokeapi.PokeApiClient
	pokemonsCaught map[string]model.Pokemon
}

func NewCatchCommand(client *pokeapi.PokeApiClient, pokemonsCaught map[string]model.Pokemon) catchCommand {
	return catchCommand{pokeApiClient: client, pokemonsCaught: pokemonsCaught}
}

func (c catchCommand) Name() string {
	return "catch"
}

func (c catchCommand) Description() string {
	return "Tries to catch a given pokemon"
}

func (c catchCommand) Execute(args []string) error {
	if args == nil || len(args) != 1 {
		return fmt.Errorf("correct usage: 'catch <pokemon_name>'")
	}
	pokemonName := args[0]
	pokemonInfo, err := c.pokeApiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchSuccessful := tryCatch(pokemonInfo.BaseExperience)

	if catchSuccessful {
		fmt.Println(pokemonName, "was caught!")
		c.pokemonsCaught[pokemonName] = pokemonInfo
		return nil
	}
	fmt.Println(pokemonName, "escaped!")
	return nil
}

func tryCatch(pokemonBaseExp int) bool {
	return rand.IntN(280) > pokemonBaseExp
}
