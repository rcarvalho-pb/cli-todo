package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/cli-todo/pkg/db"
)

type Config struct {
	DDL     string
	dbPath  string
	DB      *sql.DB
	Queries *db.Queries
}

func GetConfig(ddl string) *Config {
	dbPath := os.Getenv("DB_PATH")
	return &Config{
		DDL:    ddl,
		dbPath: dbPath,
	}
}

func (c *Config) StartConfig() error {
	if err := c.dbExists(); err != nil {
		return err
	}

	ctx := context.Background()
	conn := c.initDB()

	if _, err := conn.ExecContext(ctx, c.DDL); err != nil {
		return err
	}

	c.Queries = db.New(conn)

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
		path := strings.Split(c.dbPath, string(os.PathSeparator))
		path = path[:2]
		err = os.MkdirAll(filepath.Join(path...), os.ModePerm)
		if err != nil {
			return err
		}

		_, err = os.Create(c.dbPath)

		log.Println("DB created!")
	}
	return nil
}

func (c *Config) Ending() {
	defer c.DB.Close()
}
