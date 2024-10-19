package config

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DDL    string
	dbPath string
}

func GetConfig(ddl string) *Config {
	dbPath := os.Getenv("DB_PATH")
	return &Config{
		DDL:    ddl,
		dbPath: dbPath,
	}
}

func (c *Config) StartConfig() error {

	if _, err := os.Stat(c.dbPath); err != nil {
		err = os.MkdirAll(c.dbPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	ctx := context.Background()
	db := c.connectToDB()
	return nil
}

// func (c *Config) connectToDB() sql.DB {
//     conn :=
// }

func (c *Config) openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", c.dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
