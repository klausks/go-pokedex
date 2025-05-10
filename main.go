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
	exploreCommand := cli.NewExploreCommand(pokeApiClient)
	catchCommand := cli.NewCatchCommand(pokeApiClient)
	helpCommand := cli.NewHelpCommand([]cli.CliCommand{mapCommand, mapbCommand, exitCommand, exploreCommand, catchCommand})

	return map[string]cli.CliCommand{
		mapCommand.Name():     mapCommand,
		mapbCommand.Name():    mapbCommand,
		exitCommand.Name():    exitCommand,
		exploreCommand.Name(): exploreCommand,
		helpCommand.Name():    helpCommand,
		catchCommand.Name():   catchCommand,
	}
}

func main() {
	availableCommands := initAvailableCommands()

	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		commandLine := cleanInput(inputScanner.Text())
		commandStr := commandLine[0]
		var args []string
		if len(commandLine) > 1 {
			args = commandLine[1:]
		}

		if command, exists := availableCommands[commandStr]; exists {
			err := command.Execute(args)
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
