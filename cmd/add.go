package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dungtc/go-cli-playground/db"
)

// NewAddCmd creates add command
func NewAddCmd(task *db.TaskRepository) *Cmd {
	return &Cmd{
		task:        task,
		cmd:         flag.NewFlagSet("add", flag.ExitOnError),
		Description: "add Add a new task",
	}
}

// CreateTask creates a new task
func CreateTask(add *Cmd) {
	// addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	add.cmd.Parse(os.Args[2:])
	msg := strings.Join(os.Args[2:], "")
	fmt.Printf("New task: %v\n", msg)

	_, err := add.task.CreateTask(msg)
	if err != nil {
		os.Exit(1)
	}
}
