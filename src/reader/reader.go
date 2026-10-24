package reader

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Graph struct {
	Nodes       []Node       `yaml:"nodes"`
	Connections []Connection `yaml:"connections"`
}

type Node struct {
	Name string `yaml:"name"`
}

type Connection struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

func Read(fileName string) (Graph, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Graph yaml could not be readed")
	}

	var graph Graph
	err = yaml.Unmarshal(data, &graph)
	if err != nil {
		log.Fatalf("Error unmarshalling graph")
	}

	return graph, nil
}
