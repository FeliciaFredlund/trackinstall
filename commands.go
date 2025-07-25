package main

import (
	"fmt"
	"os"
	"strings"
)

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"add": {
			name:        "Add",
			description: "Add a new program you are about to install or have just installed",
			callback:    commandAdd,
		},
		"help": {
			name:        "Help",
			description: "Usage information and available commands.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "Exit",
			description: "Saves changes and exits the program.",
			callback:    commandExit,
		},
	}
}

type replCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

func commandExit(config *config) error {
	fmt.Println("Saving your changes to file...")
	// CALL FUNCTION TO SAVE ALL DATA TO FILE
	fmt.Println("Done.")
	os.Exit(0)
	return fmt.Errorf("something went wrong while closing the program")
}

func commandHelp(config *config) error {
	fmt.Println("\nThis is a tool to track manually installed programs.")
	fmt.Print("\nAvailable commands:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandAdd(config *config) error {
	fmt.Println("Adding a program")
	fmt.Println("A program has a name, dependencies (the packages you need to install with your package manager), installation instructions, and uninstall instructions.")
	fmt.Println("Name is required. Dependencies and the instructions are optional.")
	fmt.Println()

	fmt.Print("Name: ")
	config.reader.Scan()
	name := config.reader.Text()
	if name == "" {
		return fmt.Errorf("name is required")
	}

	fmt.Print("Dependencies (names of packages with space as separator): ")
	config.reader.Scan()
	packageNames := strings.Fields(config.reader.Text())

	fmt.Print("Installation instructions: TO BE IMPLEMENTED\n")
	// Read a file option (install.sh script or similar) --- input steps one by one --- input all steps at once?

	fmt.Print("Uninstall instructions: TO BE IMPLEMENTED\n")
	// Read a file option (install.sh script or similar) --- input steps one by one --- input all steps at once?

	if _, exists := config.programs[name]; exists {
		return fmt.Errorf("program already exists, pick a different name or edit existing program")
	}

	prog := newProgram(name, packageNames, config.dependencies)
	config.programs[name] = &prog

	return nil
}
