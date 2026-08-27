package cmd

import "codeberg.org/OliveiraJ/pokedexcli/internal/domain"

func GetCommands() map[string]domain.CliCommand {
	return map[string]domain.CliCommand{
		ExitName:    ExitCommand(),
		HelpName:    HelpCommand(),
		MapName:     MapCommand(),
		MapBName:    MapBCommand(),
		ExploreName: ExploreCommand(),
		CatchName:   CatchCommand(),
		InspectName: InspectCommand(),
		PokedexName: PokedexCommand(),
	}
}
