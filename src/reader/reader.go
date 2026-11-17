package reader

import (
	"fmt"
	"math/rand/v2"
	"nwgraph/entities"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
	NodeRadius   = 20
)

type rawGraph struct {
	Nodes       []*entities.Node `yaml:"nodes"`
	Connections []rawConnection  `yaml:"connections"`
}

type rawConnection struct {
	From   string  `yaml:"from"`
	To     string  `yaml:"to"`
	Weight float32 `yaml:"weight"`
}

func Read(fileName string) (*entities.Graph, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var raw rawGraph
	err = yaml.Unmarshal(data, &raw)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling graph: %v", err)
	}

	nodeMap := make(map[string]*entities.Node)
	for _, node := range raw.Nodes {
		node.X = float32(rand.IntN(ScreenWidth - 100))
		node.Y = float32(rand.IntN(ScreenHeight - 100))
		nodeMap[node.Name] = node
	}

	var connections []*entities.Connection
	for _, rc := range raw.Connections {
		fromNode := nodeMap[rc.From]
		toNode := nodeMap[rc.To]

		if fromNode == nil || toNode == nil {
			return nil, fmt.Errorf("invalid connection: node %s or %s not found", rc.From, rc.To)
		}

		connections = append(connections, &entities.Connection{
			From:   fromNode,
			To:     toNode,
			Weight: rc.Weight,
		})
	}

	graph := &entities.Graph{
		Nodes:       raw.Nodes,
		Connections: connections,
	}

	return graph, nil
}
