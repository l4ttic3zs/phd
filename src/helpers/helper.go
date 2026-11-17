package helpers

import (
	"nwgraph/entities"
	"slices"
)

func AddNodeIfMissing(nodes []entities.Node, newNode entities.Node) []entities.Node {
	exists := slices.ContainsFunc(nodes, func(n entities.Node) bool {
		return n.ID == newNode.ID
	})

	if !exists {
		return append(nodes, newNode)
	}
	return nodes
}
