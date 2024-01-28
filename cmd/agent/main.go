package main

import (
	"os"

	"github.com/levidousseaux/plataform-agent/internal/entity"
	"github.com/levidousseaux/plataform-agent/internal/usecases"
)

func main() {
	yamlFile, err := os.ReadFile("scripts/spa/example.yaml")
	if err != nil {
		panic(err)
	}

	definition, err := entity.NewDefinition(yamlFile)
	if err != nil {
		panic(err)
	}

	err = usecases.RunPipeline(definition)

	if err != nil {
		panic(err)
	}
}
