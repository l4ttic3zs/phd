package main

import (
	"flag"
	"fmt"
	"log"
	"nwgraph/entities"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	file := flag.String("graph", "", "Graph YAML path to file")
	flag.Parse()
	if *file == "" {
		fmt.Println("Graph file location is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ebiten.SetWindowSize(entities.ScreenWidth, entities.ScreenHeight)
	ebiten.SetWindowTitle("")

	graph := entities.NewGraph(*file)
	if err := ebiten.RunGame(graph); err != nil {
		log.Fatal(err)
	}
}
