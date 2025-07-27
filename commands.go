package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type replCommand struct {
	name        string
	description string
	callback    func(config *config) error
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
	fmt.Print("\nAvailable commands (optionally add the program name after the command):\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandAdd(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Adding a program ~*~*~")
	fmt.Println()
	fmt.Println("A program has a name and dependencies (the packages you need to install with your package manager).")
	fmt.Println("Name is required and dependencies are optional.")
	fmt.Println()

	name := getNewProgramName(config)

	fmt.Print("Dependencies (names of packages with space as separator): ")
	config.reader.Scan()
	packageNames := strings.Fields(config.reader.Text())

	prog := newProgram(name, packageNames, config.dependencies)
	config.programs[strings.ToLower(name)] = &prog

	return nil
}

func commandEdit(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Editing a program ~*~*~")
	fmt.Println()
	fmt.Println("1. Get the name of the program you want to edit.")
	fmt.Println("2. Editing begins.")
	fmt.Println()

	name := getExistingProgramName(config)
	prog := config.programs[name]
	oldName := prog.name

	fmt.Println("\nCurrent program data:")
	fmt.Println()
	fmt.Println(prog)

	for editing := true; editing; {
		fmt.Println()
		fmt.Println("What would you like to edit? Answer with one number (1-3). It will loop so you can pick something else next.")
		fmt.Println("1 - Name")
		fmt.Println("2 - Dependencies")
		fmt.Println("3 - Finished editing")

		fmt.Print("Answer: ")
		config.reader.Scan()
		answer := config.reader.Text()

		fmt.Println()

		switch answer {
		case "1": // CHANGE NAME
			fmt.Println("Current name:", prog.name)
			fmt.Println()

			config.programName = "" // resetting the name so it doesn't prefill it when editing the name
			newName := getNewProgramName(config)

			delete(config.programs, strings.ToLower(prog.name))
			config.programs[strings.ToLower(newName)] = prog

			for _, dep := range prog.dependencies {
				i := slices.Index(dep.programs, prog.name)
				dep.programs[i] = newName
			}

			prog.name = newName

		case "2": // CHANGE DEPENDENCIES
			fmt.Println("Current dependencies:", programDepedenciesString(prog.dependencies))
			fmt.Println()

			fmt.Print("Write the names of the ones you want to add: ")
			config.reader.Scan()
			add := strings.Fields(config.reader.Text())

			fmt.Println()
			fmt.Print("Write the names of the ones you want to remove: ")
			config.reader.Scan()
			remove := strings.Fields(config.reader.Text())

			obsoleteDependencies := changeDependencies(prog, add, remove, config.dependencies)

			if len(obsoleteDependencies) != 0 {
				printObsoleteDepedencies(obsoleteDependencies)
			}

		case "3": // FINISHED EDITING
			editing = false
			fmt.Println("Editing finished.")

		default: // BLANK ANSWER, ESCAPE OR INVALID CHOICE
			fmt.Println("Invalid choice.")
		}
	}

	fmt.Println()
	fmt.Printf("Updating %s to:\n %s\n", oldName, prog)

	return nil
}

func commandRemove(config *config) error {
	fmt.Println()
	fmt.Println("~*~*~ Removing a program ~*~*~")
	fmt.Println()
	fmt.Println("1. Get name of the program you wish to delete.")
	fmt.Println("2. Confirm you want to delete the program.")
	fmt.Println("3. Delete the program.")
	fmt.Println()

	name := getExistingProgramName(config)
	prog := config.programs[name]

	fmt.Println()
	fmt.Printf("Do you want to delete \"%s\"?\nY/N: ", prog.name)
	config.reader.Scan()
	answer := strings.ToLower(config.reader.Text())
	fmt.Println()

	if strings.HasPrefix(strings.ToLower(answer), "y") {
		fmt.Printf("Begin deleting %s...\n", prog.name)

		progDeps := []string{}
		for _, dep := range prog.dependencies {
			progDeps = append(progDeps, dep.name)
		}

		obsoleteDependencies := changeDependencies(prog, nil, progDeps, config.dependencies)
		if len(obsoleteDependencies) != 0 {
			printObsoleteDepedencies(obsoleteDependencies)
		}

		delete(config.programs, name)

		fmt.Println()
		fmt.Println("Deleting done.")
	} else {
		fmt.Println("Will not delete:", prog.name)
	}

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

	fmt.Printf("\nHere is the info on %s:\n", prog.name)
	fmt.Printf("Dependencies: %s", programDepedenciesString(prog.dependencies))
	fmt.Println()

	return nil
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
		} else if _, exists := config.programs[strings.ToLower(name)]; exists {
			fmt.Println("Program already exists. Try again")
			name = ""
		} else {
			return name
		}
	}
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

func printObsoleteDepedencies(obsoleteDeps []*dependency) {
	fmt.Println()
	fmt.Println("~*~*~")
	fmt.Println("OBS! There are dependencies no longer used by any tracked program.")
	fmt.Println("It is safe to set them to auto in your package manager.")
	fmt.Println()
	fmt.Println("Obsolete dependencies:")
	fmt.Println(programDepedenciesString(obsoleteDeps))
	fmt.Println("~*~*~")
}
