package entities

import (
	"errors"
	"image/color"
	"log"
	"math"
	"nwgraph/reader"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	NodeRadius   = 20
)

type Graph struct {
	nodes       []*Node
	Connections []*Connection
	draggedNode *Node
	dragOffsetX float32
	dragOffsetY float32
}

func NewGraph(file string) *Graph {

	rGraph, err := reader.Read(file)
	if err != nil {
		log.Fatalf("Error inicializing Graph")
	}
	var graph Graph
	var nodes []*Node
	var connections []*Connection

	for index, node := range rGraph.Nodes {
		nodes = append(nodes, &Node{ID: index, Type: node.Name, X: ScreenWidth / float32(len(rGraph.Nodes)) * float32(index), Y: 200})
	}
	graph.nodes = nodes

	for _, connection := range rGraph.Connections {
		from, err := graph.getNode(connection.From)
		if err != nil {
			log.Fatalf("Error finding node From")
		}
		to, err := graph.getNode(connection.To)
		if err != nil {
			log.Fatalf("Error finding node To")
		}
		connections = append(connections, &Connection{From: from, To: to})
	}
	graph.Connections = connections

	return &graph
}

func (g *Graph) Update() error {
	mx, my := ebiten.CursorPosition()
	curX, curY := float32(mx), float32(my)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, node := range g.nodes {
			dx := curX - node.X
			dy := curY - node.Y

			if math.Hypot(float64(dx), float64(dy)) <= NodeRadius {
				g.draggedNode = node
				g.dragOffsetX = node.X - curX
				g.dragOffsetY = node.Y - curY
				break
			}
		}
	}

	if g.draggedNode != nil {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.draggedNode.X = curX + g.dragOffsetX
			g.draggedNode.Y = curY + g.dragOffsetY
		} else {

			g.draggedNode = nil
		}
	}

	return nil
}

func (g *Graph) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 30, G: 30, B: 35, A: 255})

	ConnectionColor := color.RGBA{R: 120, G: 140, B: 160, A: 255}
	for _, Connection := range g.Connections {
		vector.StrokeLine(
			screen,
			Connection.From.X, Connection.From.Y,
			Connection.To.X, Connection.To.Y,
			2.5,
			ConnectionColor,
			true,
		)
	}

	for _, node := range g.nodes {
		nodeColor := color.RGBA{R: 50, G: 150, B: 250, A: 255}
		if node == g.draggedNode {
			nodeColor = color.RGBA{R: 255, G: 100, B: 100, A: 255}
		}

		vector.DrawFilledCircle(screen, node.X, node.Y, NodeRadius, nodeColor, true)

		vector.StrokeCircle(screen, node.X, node.Y, NodeRadius, 2, color.White, true)

		labelX := int(node.X) - (len(node.Type) * 3)
		labelY := int(node.Y) + NodeRadius + 5
		ebitenutil.DebugPrintAt(screen, node.Type, labelX, labelY)
	}

	ebitenutil.DebugPrint(screen, "")
}

func (g *Graph) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Graph) getNode(name string) (*Node, error) {
	if g.nodes == nil {
		err := errors.New("Empty Nodes")
		return &Node{}, err
	}
	for _, node := range g.nodes {
		if name == node.Type {
			return node, nil
		}
	}
	return &Node{}, errors.New("Node not found")
}
