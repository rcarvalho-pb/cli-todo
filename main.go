package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/rcarvalho-pb/cli-todo/internal/config"
)

//go:embed db/sqlc/schema.sql
var ddl string

func main() {

	config := config.GetConfig(ddl)

	config.StartConfig()

	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int64("complete", 1, "finish a task")

	flag.Parsed()

	switch {
	case *add:
		config.Models.Task.AddTask("test task")
	case *complete > 0:
		config.Models.Task.ToggleTask(*complete)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}

}
