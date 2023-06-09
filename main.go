package main

import (
	"log"

	"github.com/nixoncode/go_todo/cmd"
	"github.com/nixoncode/go_todo/config"
	"github.com/spf13/cobra"
)

func init() {
	err := config.LoadENV()
	if err != nil {
		log.Fatalln("Failed to load env", err)
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatalln("Failed to open DB connection", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to DB", err)
	}
}

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
