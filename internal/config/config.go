package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/cli-todo/internal/models"
	"github.com/rcarvalho-pb/cli-todo/pkg/db"
)

type Config struct {
	DDL     string
	dbPath  string
	DB      *sql.DB
	Queries *db.Queries
	Models  *models.Models
}

func GetConfig(ddl string) *Config {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		DDL:    ddl,
		dbPath: filepath.Join(dir, "db.db"),
	}
}

func (c *Config) StartConfig() error {
	if err := c.dbExists(); err != nil {
		return err
	}

	ctx := context.Background()
	c.DB = c.initDB()

	if _, err := c.DB.ExecContext(ctx, c.DDL); err != nil {
		return err
	}

	c.Queries = db.New(c.DB)
	c.Models = models.NewModels(c.Queries)

	return nil
}

func (c *Config) initDB() *sql.DB {
	db := c.connectToDB()
	if db == nil {
		log.Fatal("Couldn't connect to the DB")
	}

	return db
}

func (c *Config) connectToDB() *sql.DB {
	count := 0
	for {
		if count == 10 {
			return nil
		}
		conn, err := c.openDB()
		if err != nil {
			log.Println("DB not ready yet...")
		} else {
			return conn
		}
		count++
		fmt.Println("backing off for 1 sec...")
		time.Sleep(1 * time.Second)
	}
}

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

func (c *Config) dbExists() error {
	if _, err := os.Stat(c.dbPath); err != nil {
		log.Println("DB not found, creating db...")
		_, err = os.Create(c.dbPath)

		log.Println("DB created!")
	}
	return nil
}

func (c *Config) Ending() {
	defer c.DB.Close()
}
