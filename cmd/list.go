package cmd

import (
	"flag"
	"os"

	"github.com/dungtc/go-cli-playground/db"
)

// NewAddCmd creates add command
func NewListCmd(task *db.TaskRepository) *Cmd {
	return &Cmd{
		task:        task,
		cmd:         flag.NewFlagSet("list", flag.ExitOnError),
		Description: "list List of tasks",
	}
}

// ListTask gets list of tasks
func ListTask(list *Cmd) {
	list.cmd.Parse(os.Args[2:])

	_, err := list.task.ListTask()
	if err != nil {
		os.Exit(1)
	}
}
