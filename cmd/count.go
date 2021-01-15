package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dungtc/go-cli-playground/db"
)

// NewCountCmd creates add command
func NewCountCmd(task *db.TaskRepository) *Cmd {
	return &Cmd{
		task:        task,
		cmd:         flag.NewFlagSet("count", flag.ExitOnError),
		Description: "count Count total tasks",
	}
}

// Count get total key/pairs tasks
func Count(countCmd *Cmd) {
	countCmd.cmd.Parse(os.Args[2:])

	count, err := countCmd.task.Count()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Count: %v\n", count)
}
