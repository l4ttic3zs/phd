package main

import (
	"fmt"
	"nwgraph/commands"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var graphFile string

	var rootCmd = &cobra.Command{
		Use:   "nwgraph",
		Short: "",
	}

	allcommand := commands.GetCommands(graphFile)

	for _, c := range allcommand {
		rootCmd.AddCommand(&c)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
