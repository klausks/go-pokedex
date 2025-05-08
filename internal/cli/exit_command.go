package cli

import (
	"fmt"
	"os"
)

type exitCommand struct{}

func NewExitCommand() CliCommand {
	return exitCommand{}
}

func (c exitCommand) Name() string {
	return "exit"
}

func (c exitCommand) Description() string {
	return "Exit the Pokedex"
}

func (c exitCommand) Execute(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
