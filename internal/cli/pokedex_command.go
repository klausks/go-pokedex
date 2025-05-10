package cli

import (
	"fmt"
	"strings"

	"github.com/klausks/go-pokedex/model"
)

type pokedexCommand struct {
	pokemonsCaught map[string]model.Pokemon
}

func NewpokedexCommand(pokemonsCaught map[string]model.Pokemon) pokedexCommand {
	return pokedexCommand{pokemonsCaught: pokemonsCaught}
}

func (c pokedexCommand) Name() string {
	return "pokedex"
}

func (c pokedexCommand) Description() string {
	return "Displays a list of pokemons caught"
}

func (c pokedexCommand) Execute(args []string) error {
	caughtPokemonNames := make([]string, len(c.pokemonsCaught))
	idx := 0
	for name,_ := range c.pokemonsCaught {
		caughtPokemonNames[idx] = name
		idx++
	}
	caughtPomemonNamesList := strings.Join(caughtPokemonNames, "\n-")

	fmt.Printf("Your pokedex:\n-%s\n", caughtPomemonNamesList)
	return nil
}
