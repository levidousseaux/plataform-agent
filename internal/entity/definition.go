package entity

import (
	"errors"

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
	Name      string            `yaml:"name"`
	Template  string            `yaml:"template,omitempty"`
	Commands  []string          `yaml:"commands,omitempty"`
	Variables map[string]string `yaml:"variables,omitempty"`
}

func NewDefinition(yamlContent []byte) (*Definition, error) {
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

	for i, stage := range def.Stages {
		for i2, step := range stage.Steps {
			if step.Template != "" {
				template, err := def.GetTemplate(&step)
				if err != nil {
					return nil, err
				}

				def.Stages[i].Steps[i2] = template
			}
		}
	}

	return &def, nil
}

func (d *Definition) GetTemplate(actualStep *Step) (Step, error) {
	if actualStep.Template == "zip" {
		return d.GetZipTemplate(actualStep)
	}

	return Step{}, errors.New("template not found")
}

func (d *Definition) GetZipTemplate(actualStep *Step) (Step, error) {
	artifactDirectory := actualStep.Variables["artifact_directory"]
	if artifactDirectory == "" {
		return Step{}, errors.New("artifact_directory not found")
	}

	zipFileName := actualStep.Variables["artifact_name"]
	if zipFileName == "" {
		return Step{}, errors.New("artifact_name not found")
	}

	commands := []string{
		"mkdir -p artifacts",
		"zip -r artifacts/" + zipFileName + ".zip " + artifactDirectory,
	}

	return Step{
		Name:      actualStep.Name,
		Commands:  commands,
		Template:  "",
		Variables: nil,
	}, nil
}
