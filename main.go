package main

import (
	_ "embed"
	"fmt"

	"github.com/rcarvalho-pb/cli-todo/internal/config"
)

//go:embed db/sqlc/schema.sql
var ddl string

func main() {

	fmt.Println("Starting program")
	config := config.GetConfig(ddl)

	config.StartConfig()
	fmt.Println("Ending program")
}
