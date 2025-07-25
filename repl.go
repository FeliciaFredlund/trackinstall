package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replLoop() {
	reader := bufio.NewScanner(os.Stdin)
	config := config{
		reader:       reader,
		programs:     map[string]*program{},
		dependencies: map[string]*dependency{},
	}
	commands := getCommands()

	for {
		fmt.Print("TrackInstall > ")
		reader.Scan()

		input := cleanInput(reader.Text())

		if len(input) == 0 {
			continue
		}

		if command, exists := commands[input[0]]; exists {
			if err := command.callback(&config); err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		/*for key, prog := range config.programs {
			fmt.Print(key, prog.name, " ")
			for _, dep := range prog.dependencies {
				fmt.Print(dep.name, dep.programs, " ")
			}
			fmt.Println()
		}*/
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}

type config struct {
	reader       *bufio.Scanner
	programs     map[string]*program
	dependencies map[string]*dependency
}
