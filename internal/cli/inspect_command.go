package cli

import (
	"fmt"

	"github.com/klausks/go-pokedex/model"
)

type inspectCommand struct {
	pokemonsCaught map[string]model.Pokemon
}

func NewinspectCommand(pokemonsCaught map[string]model.Pokemon) inspectCommand {
	return inspectCommand{pokemonsCaught: pokemonsCaught}
}

func (c inspectCommand) Name() string {
	return "inspect"
}

func (c inspectCommand) Description() string {
	return "Shows information about a caught pokemon"
}

func (c inspectCommand) Execute(args []string) error {
	if args == nil || len(args) != 1 {
		return fmt.Errorf("correct usage: 'inspect <pokemon_name>'")
	}
	pokemonName := args[0]
	pokemon, exists := c.pokemonsCaught[pokemonName]
	if !exists {
		return fmt.Errorf("could not inspect requested pokemon because it was not caught yet")
	}

	fmt.Println(pokemon.Info())
	return nil
}
