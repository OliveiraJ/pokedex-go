package cmd

import (
	"encoding/json"
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	MapBName        = "mapb"
	MapBDescription = "List all available maps"
)

type MapBCommandGetResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func MapBCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        MapBName,
		Description: MapBDescription,
		Callback:    commandMapB,
	}
}

func commandMapB(config *domain.Config, args []string) error {
	if config.Previous == "" {
		return fmt.Errorf("no previous map")
	}

	var locationAreas MapCommandGetResponse
	body, err := config.PokeAPIClient.Get(config.Previous)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	config.Next = locationAreas.Next
	if locationAreas.Previous != nil {
		config.Previous = locationAreas.Previous.(string)
	}

	for _, m := range locationAreas.Results {
		fmt.Println(m.Name)
	}

	return nil
}
