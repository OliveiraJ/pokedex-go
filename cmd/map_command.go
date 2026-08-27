package cmd

import (
	"encoding/json"
	"fmt"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

const (
	MapName        = "map"
	MapDescription = "List all available maps"
)

type MapCommandGetResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func MapCommand() domain.CliCommand {
	return domain.CliCommand{
		Name:        MapName,
		Description: MapDescription,
		Callback:    commandMap,
	}
}

func commandMap(config *domain.Config, args []string) error {
	var locationAreas MapCommandGetResponse
	body, err := config.PokeAPIClient.Get(config.Next)
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
