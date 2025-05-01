package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/klausks/go-pokedex/internal"
)

type cliCommand struct {
	name        string
	description string
	context     *commandContext
	callback    func(*commandContext) error
}

type commandContext struct {
	previous string
	next     string
}

var availableCommands map[string]cliCommand

func initAvailableCommands() {
	availableCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			context:     &commandContext{},
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the location areas",
			context:     &commandContext{},
			callback:    commandMap,
		},
	}

	// Now add help command after map exists
	availableCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

func main() {
	initAvailableCommands()
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		commandStr := cleanInput(inputScanner.Text())[0]

		if command, exists := availableCommands[commandStr]; exists {
			err := command.callback(command.context)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Uknown command")
		}
	}
}

func commandExit(context *commandContext) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(context *commandContext) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println(getAvailableCommandsHelp())
	return nil
}

func commandMap(context *commandContext) error {
	resp, err := internal.GetLocationAreaNames(context.next)
	if err != nil {
		return err
	}

	locationAreaNames := getLocationAreaNames(resp)
	for _, areaName := range locationAreaNames {
		fmt.Println(areaName)
	}

	return nil
}

func getLocationAreaNames(resp internal.LocationAreaBatch) []string {
	var locationAreaNames = make([]string, len(resp.Results))
	for i, locationArea := range resp.Results {
		// locationAreaNames = append(locationAreaNames, locationArea.Name)
		locationAreaNames[i] = locationArea.Name
	}
	return locationAreaNames
}

func getAvailableCommandsHelp() string {
	var commandHelpStrings []string
	for _, command := range availableCommands {
		commandHelpString := fmt.Sprintf("%s: %s", command.name, command.description)
		commandHelpStrings = append(commandHelpStrings, commandHelpString)
	}
	return strings.Join(commandHelpStrings, "\n")
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
}
