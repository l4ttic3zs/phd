package dijkstra

import (
	"container/heap"
	"log"
	"nwgraph/entities"
)

type DijkstraResult struct {
	Path        []*entities.Node
	TotalWeight float32
}

func FindShortestPath(start, target *entities.Node) DijkstraResult {
	pqItems := make(map[int]*Item)

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Kezdeti beállítások (végtelen távolság mindenhol, kivéve a start)
	// Megjegyzés: mivel a Node ID-t használjuk kulcsként, feltételezzük, hogy van egyedi ID-ja.
	// Ha a pointert magát használjuk map kulcsként, az is teljesen biztonságos Go-ban:

	// Használhatunk pointereket is a map kulcsokhoz a biztonság kedvéért:
	distMap := make(map[*entities.Node]float32)
	prevMap := make(map[*entities.Node]*entities.Node)

	// Feltételezzük, hogy van egy globális listád vagy elérhető összes node a hálózatban,
	// de a Dijkstra dinamikusan is felfedezheti őket a Connections-ökön keresztül.

	distMap[start] = 0

	startItem := &Item{node: start, distance: 0}
	heap.Push(pq, startItem)
	pqItems[start.ID] = startItem

	for pq.Len() > 0 {
		currentIt := heap.Pop(pq).(*Item)
		currentNode := currentIt.node

		// Ha elértuk a célt, megállhatunk (vagy folytathatjuk, ha az összeset keressük)
		if currentNode == target {
			break
		}

		// Ha az aktuális távolság rosszabb, mint amit már feljegyeztünk, ugorjuk át
		if currentIt.distance > distMap[currentNode] {
			continue
		}

		// Végigmegyünk a node csatlakozásain (Connections)
		for _, conn := range currentNode.Connections {
			// Megállapítjuk, hogy melyik a szomszéd node a kapcsolat másik oldala
			var neighbor *entities.Node
			if conn.From == currentNode {
				neighbor = conn.To
			} else if conn.To == currentNode {
				neighbor = conn.From
			} else {
				continue // Biztonsági ellenőrzés
			}

			// Új kalkulált távolság (Edgemetrika / Weight hozzáadása)
			newDist := distMap[currentNode] + conn.Weight
			log.Println(conn.Weight)

			// Ha még nincsen benne a térképen, vagy találtunk egy olcsóbb (kisebb súlyú) utat:
			if d, exists := distMap[neighbor]; !exists || newDist < d {
				distMap[neighbor] = newDist
				prevMap[neighbor] = currentNode

				if item, ok := pqItems[neighbor.ID]; ok {
					pq.update(item, neighbor, newDist)
				} else {
					newItem := &Item{node: neighbor, distance: newDist}
					heap.Push(pq, newItem)
					pqItems[neighbor.ID] = newItem
				}
			}
		}
	}

	// Útvonal visszafejtése a target-től a start felé
	var path []*entities.Node
	curr := target
	// Ellenőrizzük, hogy elértük-e egyáltalán a célt
	if _, exists := distMap[target]; !exists && target != start {
		return DijkstraResult{Path: nil, TotalWeight: -1} // Nincs út
	}

	for curr != nil {
		path = append([]*entities.Node{curr}, path...) // Előre beszúrás, hogy sorrendben legyen
		curr = prevMap[curr]
	}

	return DijkstraResult{
		Path:        path,
		TotalWeight: distMap[target],
	}
}
