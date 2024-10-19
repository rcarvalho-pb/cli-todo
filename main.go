package main

import (
	_ "embed"

	"github.com/rcarvalho-pb/cli-todo/internal/config"
)

//go:embed db/migrations/000001_create_tables.up.sql
var ddl string

func main() {

	config := config.GetConfig(ddl)

	config.StartConfig()
}
