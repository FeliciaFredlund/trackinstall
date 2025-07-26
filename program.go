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

func (p program) String() string {
	text := "\"" + p.name + "\" depends on "
	for _, dep := range p.dependencies {
		text += "\"" + dep.name + "\" "
	}
	if len(p.dependencies) == 0 {
		text += "n/a"
	}
	//text = strings.TrimSuffix(text, ", ")
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
