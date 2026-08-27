package cmd

import (
	"fmt"
	"os"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	ExitName        = "exit"
	ExitDescription = "Exit the application"
)

func ExitCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        ExitName,
		Description: ExitDescription,
		Callback:    commandExit,
	}
}

func commandExit(config *domain.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
