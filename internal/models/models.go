package models

import (
	"time"

	"github.com/rcarvalho-pb/cli-todo/pkg/db"
)

var queries *db.Queries

const dbTimeout = 10 * time.Second

type Models struct {
	Task *Task
}

func NewModels(dbPool *db.Queries) *Models {
	queries = dbPool

	return &Models{
		Task: &Task{},
	}
}
