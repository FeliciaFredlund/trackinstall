package main

import (
	"fmt"
	"os"
	"strings"
)

type replCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

func getCommands() map[string]replCommand {
	return map[string]replCommand{
		"add": {
			name:        "Add",
			description: "Add a new program that you want to track",
			callback:    commandAdd,
		},
		"edit": {
			name:        "Edit",
			description: "Edit a program",
			callback:    commandEdit,
		},
		"remove": {
			name:        "Remove",
			description: "Remove a program and learn if you have depedencies you can uninstall",
			callback:    commandRemove,
		},
		"list": {
			name:        "List",
			description: "List all tracked programs and their dependencies",
			callback:    commandList,
		},
		"show": {
			name:        "Show",
			description: "Show all the data for one specified program",
			callback:    commandShow,
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

func commandExit(config *config) error {
	fmt.Println("Saving your changes to file...")
	err := saveData(SAVEFILEPATH, config.programs)
	if err != nil {
		fmt.Println("Something went wrong when saving data. Aborting exiting program.")
		return fmt.Errorf("could not save data on trying to exit: %w", err)
	}
	fmt.Println("Done.")
	os.Exit(0)
	return fmt.Errorf("something went wrong while closing the program")
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Help ~*~*~")
	fmt.Println()
	fmt.Println("This is a tool to track manually installed programs and their dependencies.")
	fmt.Print("\nAvailable commands: (optionally add the program name after the command)\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandAdd(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Adding a program ~*~*~")
	fmt.Println()
	fmt.Println("A program has a name, dependencies (the packages you need to install with your package manager), installation instructions, and uninstall instructions.")
	fmt.Println("Name is required. Dependencies and the instructions are optional.")
	fmt.Println()

	name := getNewProgramName(config)

	fmt.Print("Dependencies (names of packages with space as separator): ")
	config.reader.Scan()
	packageNames := strings.Fields(config.reader.Text())

	/*fmt.Print("Installation instructions: TO BE IMPLEMENTED\n")
	// Read a file option (install.sh script or similar) --- input steps one by one --- input all steps at once?

	fmt.Print("Uninstall instructions: TO BE IMPLEMENTED\n") */
	// Read a file option (install.sh script or similar) --- input steps one by one --- input all steps at once?

	prog := newProgram(name, packageNames, config.dependencies)
	config.programs[strings.ToLower(name)] = &prog

	return nil
}

func commandEdit(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Editing a program ~*~*~")
	fmt.Println()

	return nil
}

func commandRemove(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Removing a program ~*~*~")
	fmt.Println()

	return nil
}

func commandList(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ List of all programs ~*~*~")

	for _, prog := range config.programs {
		fmt.Println("\n", prog)
	}
	if len(config.programs) == 0 {
		fmt.Println("\n No programs added yet...")
	}

	return nil
}

func commandShow(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Show program ~*~*~")
	fmt.Println()

	name := getExistingProgramName(config)
	prog := config.programs[name]
	fmt.Print("Dependencies: ")
	for _, dep := range prog.dependencies {
		fmt.Print("\"" + dep.name + "\" ")
	}
	if len(prog.dependencies) == 0 {
		fmt.Print("n/a")
	}
	fmt.Println()

	return nil
}

// Function to deal with getting and checking the name for programs
func getNewProgramName(config *config) string {
	name := config.programName

	for {
		// Get name (or use name from config)
		fmt.Print("Name: ")
		if name == "" {
			config.reader.Scan()
			name = config.reader.Text()
		} else {
			fmt.Println(name)
		}

		name = strings.TrimSpace(name)

		// Checking name
		if name == "" {
			fmt.Println("Name is required. Try again.")
		} else if prog, exists := config.programs[strings.ToLower(name)]; exists {
			fmt.Print("Program already exists.\n", "Program:", prog, "\nDo you want to overwrite it (y/n)? ")
			config.reader.Scan()
			answer := config.reader.Text()
			if strings.HasPrefix(strings.ToLower(answer), "y") {
				break
			} else {
				name = ""
			}
		} else {
			break
		}
	}

	return name
}

func getExistingProgramName(config *config) string {
	name := config.programName

	for {
		// Get name (or use name from config)
		fmt.Print("Name: ")
		if name == "" {
			config.reader.Scan()
			name = config.reader.Text()
		} else {
			fmt.Println(name)
		}

		name = strings.ToLower(strings.TrimSpace(name))

		// Checking name
		if name == "" {
			fmt.Println("Name is required. Try again.")
		} else if _, exists := config.programs[name]; !exists {
			fmt.Println("Can't find program. Please check spelling (capitalization doesn't matter).")
			name = ""
		} else {
			break
		}
	}

	return name
}
