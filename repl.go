package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	reader       *bufio.Scanner
	programs     map[string]*program
	dependencies map[string]*dependency
}

const (
	SAVEFILEPATH = "ti_data.json"
)

func replLoop() {
	reader := bufio.NewScanner(os.Stdin)
	programs, depedencies := loadData(SAVEFILEPATH)

	config := config{
		reader:       reader,
		programs:     programs,
		dependencies: depedencies,
	}
	commands := getCommands()

	for {
		fmt.Print("\nTrackInstall > ")
		reader.Scan()

		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if command, exists := commands[input[0]]; exists {
			if err := command.callback(&config); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		for key, prog := range config.programs {
			fmt.Println(key)
			for _, dep := range prog.dependencies {
				fmt.Println(dep)
			}
		}
	}
}

func loadData(filepath string) (map[string]*program, map[string]*dependency) {
	data, err := readFile(filepath)
	if err != nil {
		fmt.Printf("Error while loading data: %s\n", err)
		fmt.Println("Exiting with error")
		os.Exit(1)
	}

	programs, dependencies, err := parseData(data)
	if err != nil {
		fmt.Printf("Error while parsing data: %s\n", err)
		fmt.Println("Exiting with error")
		os.Exit(1)
	}

	return programs, dependencies
}

func saveData(filepath string, programs map[string]*program) error {
	data, err := parseStructs(programs)
	if err != nil {
		return fmt.Errorf("error while parsing data to json: %w", err)
	}
	err = writeFile(filepath, data)
	if err != nil {
		return fmt.Errorf("error while saving data to file: %w", err)
	}

	return nil
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
