package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	nodeRadius   = 20
)

type Node struct {
	ID   int
	Type string
	X, Y float32
}

type Edge struct {
	From *Node
	To   *Node
}

type Game struct {
	nodes       []*Node
	edges       []Edge
	draggedNode *Node
	dragOffsetX float32
	dragOffsetY float32
}

func NewGame() *Game {

	n1 := &Node{ID: 1, Type: "Router", X: 150, Y: 200}
	n2 := &Node{ID: 2, Type: "Switch", X: 400, Y: 150}
	n3 := &Node{ID: 3, Type: "Szerver 1", X: 650, Y: 100}
	n4 := &Node{ID: 4, Type: "Szerver 2", X: 650, Y: 300}
	n5 := &Node{ID: 5, Type: "Kliens PC", X: 400, Y: 400}

	nodes := []*Node{n1, n2, n3, n4, n5}

	edges := []Edge{
		{From: n1, To: n2},
		{From: n2, To: n3},
		{From: n2, To: n4},
		{From: n1, To: n5},
	}

	return &Game{
		nodes: nodes,
		edges: edges,
	}
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	curX, curY := float32(mx), float32(my)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, node := range g.nodes {
			dx := curX - node.X
			dy := curY - node.Y

			if math.Hypot(float64(dx), float64(dy)) <= nodeRadius {
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

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 30, G: 30, B: 35, A: 255})

	edgeColor := color.RGBA{R: 120, G: 140, B: 160, A: 255}
	for _, edge := range g.edges {
		vector.StrokeLine(
			screen,
			edge.From.X, edge.From.Y,
			edge.To.X, edge.To.Y,
			2.5,
			edgeColor,
			true,
		)
	}

	for _, node := range g.nodes {
		nodeColor := color.RGBA{R: 50, G: 150, B: 250, A: 255}
		if node == g.draggedNode {
			nodeColor = color.RGBA{R: 255, G: 100, B: 100, A: 255}
		}

		vector.DrawFilledCircle(screen, node.X, node.Y, nodeRadius, nodeColor, true)

		vector.StrokeCircle(screen, node.X, node.Y, nodeRadius, 2, color.White, true)

		labelX := int(node.X) - (len(node.Type) * 3)
		labelY := int(node.Y) + nodeRadius + 5
		ebitenutil.DebugPrintAt(screen, node.Type, labelX, labelY)
	}

	ebitenutil.DebugPrint(screen, "Hálózati Diagram (Ebitengine)\nHúzd a csomópontokat az egérrel!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Go Hálózati Rajzoló - Ebitengine")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
