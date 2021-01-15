package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dungtc/go-cli-playground/cmd"
	"github.com/dungtc/go-cli-playground/db"
)

func main() {
	storage, err := db.Init("tasks.db")
	if err != nil {
		panic(err)
	}

	taskRepo := db.NewTaskRepository(storage)
	addCmd := cmd.NewAddCmd(taskRepo)
	listCmd := cmd.NewListCmd(taskRepo)
	countCmd := cmd.NewCountCmd(taskRepo)
	rootCmd := []*cmd.Cmd{addCmd, listCmd, countCmd}

	flag.Usage = func() {
		var usage string
		for _, v := range rootCmd {
			usage += fmt.Sprintf("\n\t%s", v.Description)
		}
		fmt.Fprintf(os.Stderr, "Usage of %s:\n \tgo-cli-playground [command]\nAvailable Commands: %s\n", os.Args[0], usage)
		flag.PrintDefaults()
	}

	// Verify main command and sub command
	// os.Args[0] is main command
	// os.Args[1] is sub command
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	// count sub command args
	switch os.Args[1] {
	case "add":
		cmd.CreateTask(addCmd)
	case "list":
		cmd.ListTask(listCmd)
	case "count":
		cmd.Count(countCmd)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
