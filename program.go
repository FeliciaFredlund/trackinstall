package main

type program struct {
	name         string
	dependencies []*dependency
}

type dependency struct {
	name     string
	programs []string
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
		existingDependencies[dependencyName] = dep
	}

	return program{
		name:         programName,
		dependencies: dependencies,
	}
}
