package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replLoop() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("TrackInstall > ")
		reader.Scan()

		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if command, exists := commands[input[0]]; exists {
			if err := command.callback(); err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"exit": {
			name:        "Exit",
			description: "Saves changes and exits the program.",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Usage information and available commands.",
			callback:    commandHelp,
		},
	}
}

type replCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Saving your changes to file...")
	// CALL FUNCTION TO SAVE ALL DATA TO FILE
	fmt.Println("Done.")
	os.Exit(0)
	return fmt.Errorf("something went wrong while closing the program")
}

func commandHelp() error {
	fmt.Println("\nThis is a tool to track manually installed programs.")
	fmt.Print("\nAvailable commands:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
