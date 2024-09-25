package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://postgres:admin@localhost/db?sslmode=disable"
	ctx := context.Background()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected!")
}
