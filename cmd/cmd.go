package cmd

import (
	"flag"

	"github.com/dungtc/go-cli-playground/db"
)

// Cmd presents command model
type Cmd struct {
	task        *db.TaskRepository
	cmd         *flag.FlagSet
	Description string
}
