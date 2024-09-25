package storage

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // driver for sql
)

type Storage struct {
	DB *sql.DB
}

var Context context.Context

func (storage *Storage) NewDB(driver string, dsn string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	context := context.Background()
	storage.DB = db
	Context = context
	return nil
}

func (storage *Storage) AddDB(db *sql.DB) {
	Context = context.Background()
	storage.DB = db
}

func (storage *Storage) CloseDB() {
	storage.DB.Close()
}

func (storage *Storage) Ping() error {
	err := storage.DB.PingContext(Context)
	return err
}
