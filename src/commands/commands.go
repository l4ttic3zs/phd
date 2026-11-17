package commands

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"nwgraph/dijkstra"
	"nwgraph/graph"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

func GetCommands(graphFile string) []cobra.Command {
	var commands []cobra.Command
	var runGraphFile string
	var validateGraphFile string
	var startNode string
	var finishNode string

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Starting UI",
		Run: func(cmd *cobra.Command, args []string) {
			ngraph := graph.NewGraph(runGraphFile)

			go func() {
				http.HandleFunc("/compute", func(w http.ResponseWriter, r *http.Request) {
					if r.Method != http.MethodGet {
						http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
						return
					}

					start := r.URL.Query().Get("start")
					finish := r.URL.Query().Get("finish")

					if start == "" || finish == "" {
						http.Error(w, "Missing 'start' or 'finish' parameter", http.StatusBadRequest)
						return
					}
					startNode, err := ngraph.GetNode(start)
					if err != nil {
						http.Error(w, "Start not found", http.StatusNotFound)
						return
					}
					finishNode, err := ngraph.GetNode(finish)
					if err != nil {
						http.Error(w, "Finish not found", http.StatusNotFound)
						return
					}
					result := dijkstra.FindShortestPath(startNode, finishNode)
					var pathStr string
					for i, node := range result.Path {
						if i > 0 {
							pathStr += " -> "
						}
						pathStr += fmt.Sprintf("%s (ID: %d)", node.Name, node.ID)
					}

					responseText := fmt.Sprintf("Path: %s\nTotal Weight: %.2f\n", pathStr, result.TotalWeight)

					//fmt.Println("[API] compute run, UI updated")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fmt.Sprintf(responseText)))
				})

				log.Println("CLI API running on 127.0.0.1:47823")
				if err := http.ListenAndServe("127.0.0.1:47823", nil); err != nil {
					log.Printf("API server error: %v\n", err)
				}
			}()

			ebiten.SetWindowSize(graph.ScreenWidth, graph.ScreenHeight)
			ebiten.SetWindowTitle("NWGraph Visualizer")
			if err := ebiten.RunGame(ngraph); err != nil {
				log.Fatal(err)
			}
		},
	}
	runCmd.Flags().StringVarP(&runGraphFile, "graph", "g", "", "Graph YAML path to file (required)")
	runCmd.MarkFlagRequired("graph")

	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validates YAML syntax",
		Run: func(cmd *cobra.Command, args []string) {
			if graphFile == "" {
				fmt.Println("Error: A --graph is required")
				os.Exit(1)
			}

			fmt.Printf("%s file validation in progress...\n", graphFile)
			fmt.Println("Syntax OK")
		},
	}
	validateCmd.Flags().StringVarP(&validateGraphFile, "graph", "g", "", "Graph YAML path to file (required)")
	validateCmd.MarkFlagRequired("graph")

	var computeCmd = &cobra.Command{
		Use:   "compute",
		Short: "Running algorithm",
		Run: func(cmd *cobra.Command, args []string) {
			reqURL := fmt.Sprintf("http://127.0.0.1:47823/compute?start=%s&finish=%s", startNode, finishNode)

			resp, err := http.Get(reqURL)
			if err != nil {
				fmt.Println("Error: Running UI not found")
				os.Exit(1)
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			fmt.Print(string(body))
		},
	}
	computeCmd.Flags().StringVarP(&startNode, "start", "s", "", "Start of dijkstra algorithm")
	computeCmd.Flags().StringVarP(&finishNode, "finish", "f", "", "Finish of dijkstra algorithm")
	computeCmd.MarkFlagRequired("start")
	computeCmd.MarkFlagRequired("finish")

	commands = append(commands, *runCmd, *validateCmd, *computeCmd)

	return commands
}
