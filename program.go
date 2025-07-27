package main

import (
	"slices"
	"strings"
)

type program struct {
	name         string
	dependencies []*dependency
}

type dependency struct {
	name     string
	programs []string
}

func (prog program) String() string {
	text := "\"" + prog.name + "\" depends on "
	text += programDepedenciesString(prog.dependencies)
	return text
}

func programDepedenciesString(dependencies []*dependency) string {
	if len(dependencies) == 0 {
		return "n/a"
	}

	text := ""
	for _, dep := range dependencies {
		text += "\"" + dep.name + "\" "
	}
	return text
}

func newProgram(programName string, dependencyNames []string, existingDependencies map[string]*dependency) program {
	prog := program{
		name:         programName,
		dependencies: []*dependency{},
	}

	changeDependencies(&prog, dependencyNames, nil, existingDependencies)

	return prog
}

func changeDependencies(prog *program, add []string, remove []string, existingDependencies map[string]*dependency) []*dependency {
	obsoleteDependencies := []*dependency{}

	for _, depName := range add {
		dep, exists := existingDependencies[strings.ToLower(depName)]

		if exists {
			if slices.Contains(dep.programs, prog.name) { // the program already have this dependency
				continue
			}
		} else {
			dep = &dependency{
				name:     depName,
				programs: []string{},
			}
		}

		dep.programs = append(dep.programs, prog.name)
		prog.dependencies = append(prog.dependencies, dep)
		existingDependencies[strings.ToLower(depName)] = dep
	}

	//delete
	for _, depName := range remove {
		dep, exists := existingDependencies[strings.ToLower(depName)]

		if !exists { // the depedency doesn't exist
			continue
		}
		if !slices.Contains(dep.programs, prog.name) { // the program doesn't have this dependency
			continue
		}

		i := slices.Index(prog.dependencies, dep)
		prog.dependencies = slices.Delete(prog.dependencies, i, i+1)

		i = slices.Index(dep.programs, prog.name)
		dep.programs = slices.Delete(dep.programs, i, i+1)

		if len(dep.programs) == 0 {
			obsoleteDependencies = append(obsoleteDependencies, dep)
			delete(existingDependencies, strings.ToLower(dep.name))
		}
	}

	return obsoleteDependencies
}
