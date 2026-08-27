package cmd

import (
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	HelpName        = "help"
	HelpDescription = "Displays a help message"
)

func HelpCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        HelpName,
		Description: HelpDescription,
		Callback:    commandHelp,
	}
}

func commandHelp(config *domain.Config, args []string) error {
	fmt.Println(
		`Welcome to the Pokedex!
	Usage:

	help: Displays a help message
	exit: Exit the Pokedex`)
	return nil
}
