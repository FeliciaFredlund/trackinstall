package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type savedProgram struct {
	Name         string
	Dependencies []string
}

func readFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []byte{}, nil
		} else {
			return nil, fmt.Errorf("readFile with path %s failed: %w", filepath, err)
		}
	}

	return data, nil
}

func writeFile(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0664)
	if err != nil {
		return fmt.Errorf("writeFile with path %s failed: %w", filepath, err)
	}
	return nil
}

func parseData(data []byte) (map[string]*program, map[string]*dependency, error) {
	programs := map[string]*program{}
	dependencies := map[string]*dependency{}

	if len(data) == 0 { // No saved data existed, so return instanced empty maps
		return programs, dependencies, nil
	}

	savedPrograms := &[]savedProgram{}
	err := json.Unmarshal(data, savedPrograms)
	if err != nil {
		return nil, nil, fmt.Errorf("error unmarshaling saved data: %w", err)
	}

	for _, savedProg := range *savedPrograms {
		prog := newProgram(savedProg.Name, savedProg.Dependencies, dependencies)
		programs[strings.ToLower(prog.name)] = &prog
	}

	return programs, dependencies, nil
}

func parseStructs(programs map[string]*program) ([]byte, error) {
	savedPrograms := []savedProgram{}

	for _, prog := range programs {
		dependencies := make([]string, len(prog.dependencies))
		for i, dep := range prog.dependencies {
			dependencies[i] = dep.name
		}

		savedProg := savedProgram{
			Name:         prog.name,
			Dependencies: dependencies,
		}
		savedPrograms = append(savedPrograms, savedProg)
	}

	data, err := json.Marshal(savedPrograms)
	if err != nil {
		return nil, fmt.Errorf("error while converting to json: %w", err)
	}

	return data, nil
}
