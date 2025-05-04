package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/klausks/go-pokedex/internal/cli"
	"github.com/klausks/go-pokedex/internal/pokeapi"
)

func initAvailableCommands() map[string]cli.CliCommand {
	pokeApiClient := pokeapi.NewPokeApiClient()
	mapApiRequestContext := &cli.ApiRequestContext{}
	mapCommand := cli.NewMapCommand(mapApiRequestContext, pokeApiClient)
	mapbCommand := cli.NewMapbCommand(mapApiRequestContext, pokeApiClient)
	exitCommand := cli.NewExitCommand()
	helpCommand := cli.NewHelpCommand([]cli.CliCommand{mapCommand, mapbCommand, exitCommand})

	return map[string]cli.CliCommand{
		mapCommand.Name():  mapCommand,
		mapbCommand.Name(): mapbCommand,
		exitCommand.Name(): exitCommand,
		helpCommand.Name(): helpCommand,
	}
}

func main() {
	availableCommands := initAvailableCommands()
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		commandStr := cleanInput(inputScanner.Text())[0]

		if command, exists := availableCommands[commandStr]; exists {
			err := command.Execute()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Uknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}
