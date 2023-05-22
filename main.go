package main

import (
	"github.com/nixoncode/go_todo/cmd"
	"github.com/spf13/cobra"
)

func main() {
	start()
}

func start() {
	cobra := &cobra.Command{
		Use:   "todo",
		Short: "Todo CLI",
	}

	cobra.AddCommand(cmd.NewServeCommand())

	cobra.Execute()
}
