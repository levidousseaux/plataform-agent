package pipeline

import (
	"gopkg.in/yaml.v3"
)

type Definition struct {
	Image  string  `yaml:"image"`
	Stages []Stage `yaml:"stages"`
}

type Stage struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
}

func NewDefinitionFromYaml(yamlContent []byte) (*Definition, error) {
	var def Definition

	err := yaml.Unmarshal(yamlContent, &def)
	if err != nil {
		panic(err)
	}

	return &def, nil
}

func (definition *Definition) GetScript() string {
	script := ""
	for _, stage := range definition.Stages {
		script += "echo [" + stage.Name + "]\n"

		for _, actions := range stage.Steps {
			script += "echo '--" + actions.Name + "'\n"
			for _, command := range actions.Commands {
				script += command + " \n"
			}
		}
	}
	return script
}
