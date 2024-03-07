package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/mystpen/test-task-Mobydev/config"
	"github.com/pkg/errors"
)

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", "ddd")
	if err != nil {
		return nil, errors.Wrap(err, "open sql")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "connection is not established")
	}

	return db, nil
}
