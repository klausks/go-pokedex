package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var availableCommands map[string]cliCommand

func initAvailableCommands() {
	availableCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Uknown command")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println(getAvailableCommandsHelp())
	return nil
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
