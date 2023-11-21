package main

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Resources struct {
	db *sql.DB
}

func loadResources() (*Resources, error) {
	db, err := sql.Open("sqlite3", "./balancer.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	ctx := context.Background()
	q := `
		 CREATE TABLE IF NOT EXISTS servers (
            id SERIAL PRIMARY KEY ,
            host TEXT NOT NULL
        );`
	_, err = db.ExecContext(ctx, q)
	if err != nil {
		return nil, err
	}
	return &Resources{
		db: db,
	}, nil
}
