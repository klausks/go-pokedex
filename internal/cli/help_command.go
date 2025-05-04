package cli

import (
	"fmt"
	"strings"
)

type helpCommand struct {
	commandList []CliCommand
}

func NewHelpCommand(commandList []CliCommand) helpCommand {
	helpCommand := helpCommand{}
	commandList = append(commandList, helpCommand)
	helpCommand.commandList = commandList
	return helpCommand
}

func (c helpCommand) Name() string {
	return "help"
}

func (c helpCommand) Description() string {
	return "Displays help menu with information about each command"
}

func (c helpCommand) Execute() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	fmt.Println(c.getAvailableCommandsHelp())
	return nil
}

func (c helpCommand) getAvailableCommandsHelp() string {
	var commandHelpStrings []string
	for _, command := range c.commandList {
		commandHelpString := fmt.Sprintf("%s: %s", command.Name(), command.Description())
		commandHelpStrings = append(commandHelpStrings, commandHelpString)
	}
	return strings.Join(commandHelpStrings, "\n")
}
