package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		cleaned := CleanInput(text)
		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if ok == false {
			fmt.Println("Invalid Command")
			continue
		} else {
			command.callback()
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints the help menu",
			callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "lists available 20 location",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists previous 20 location",
			callback:    CommandMapB,
		},
	}
}
func CleanInput(s string) []string {
	lowered := strings.ToLower(s)
	words := strings.Fields(lowered)
	return words
}
