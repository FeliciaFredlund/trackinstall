package main

import "strings"

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
	dependencies := []*dependency{}

	for _, dependencyName := range dependencyNames {
		dep, exists := existingDependencies[dependencyName]
		if !exists {
			dep = &dependency{
				name:     dependencyName,
				programs: []string{},
			}
		}
		dep.programs = append(dep.programs, programName)
		dependencies = append(dependencies, dep)
		existingDependencies[strings.ToLower(dependencyName)] = dep
	}

	return program{
		name:         programName,
		dependencies: dependencies,
	}
}
