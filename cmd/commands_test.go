package cmd

import (
	"reflect"
	"testing"

	"codeberg.org/OliveiraJ/pokedexcli/internal/domain"
)

func TestGetCommands(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected domain.CliCommand
	}{
		{
			name:  "get help command",
			input: HelpName,
			expected: domain.CliCommand{
				Name:        HelpName,
				Description: HelpDescription,
				Callback:    commandHelp,
			},
		},
		{
			name:  "get exit command",
			input: ExitName,
			expected: domain.CliCommand{
				Name:        ExitName,
				Description: ExitDescription,
				Callback:    commandExit,
			},
		},
	}

	commands := GetCommands()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, exists := commands[c.input]
			if !exists {
				t.Fatalf("command %q was not found", c.input)
			}

			if actual.Name != c.expected.Name {
				t.Errorf("expected name %q, got %q",
					c.expected.Name, actual.Name)
			}

			if actual.Description != c.expected.Description {
				t.Errorf("expected description %q, got %q",
					c.expected.Description, actual.Description)
			}

			if actual.Callback == nil {
				t.Fatal("expected a callback, got nil")
			}

			actualCallback := reflect.ValueOf(actual.Callback).Pointer()
			expectedCallback := reflect.ValueOf(c.expected.Callback).Pointer()

			if actualCallback != expectedCallback {
				t.Error("unexpected callback")
			}
		})
	}
}
