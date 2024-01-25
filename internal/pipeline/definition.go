package pipeline

import (
	"gopkg.in/yaml.v3"
)

type Definition struct {
	Image      string  `yaml:"image"`
	Repository string  `yaml:"repository"`
	Stages     []Stage `yaml:"stages"`
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
		return nil, err
	}

	if def.Repository != "" {
		cloneStep := Step{
			Name: "Clone Repository",
			Commands: []string{
				"git clone " + def.Repository + " .",
			},
		}

		def.Stages = append([]Stage{{Name: "Setup", Steps: []Step{cloneStep}}}, def.Stages...)
	}

	return &def, nil
}
