package main

import (
	"fmt"
	"os"
	"simple_CLI_app_with_cobra_golang/tasks"
)

func main() {
	if err := tasks.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add commands and their initialization here
	tasks.AddCommands()
}
