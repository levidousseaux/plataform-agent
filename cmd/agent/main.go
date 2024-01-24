package main

import (
	"github.com/levidousseaux/plataform-agent/internal/pipeline"
	"os"
)

func main() {
	yamlFile, err := os.ReadFile("scripts/spa/example.yaml")
	if err != nil {
		panic(err)
	}

	definition, err := pipeline.NewDefinitionFromYaml(yamlFile)
	err = pipeline.RunPipeline(definition)

	if err != nil {
		panic(err)
	}
}
