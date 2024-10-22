package main

import (
	_ "embed"

	"github.com/rcarvalho-pb/cli-todo/internal/commands"
	"github.com/rcarvalho-pb/cli-todo/internal/config"
)

//go:embed db/sqlc/schema.sql
var ddl string

func main() {

	config := config.GetConfig(ddl)

	config.StartConfig()
	defer config.DB.Close()

	commands.Execute(config.Models)
}
