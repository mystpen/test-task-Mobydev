package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/mystpen/test-task-Mobydev/config"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg config.Config) (*sql.DB, error) {
	// connString := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// db, err := sql.Open("postgres", connString)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "open sql")
	// }
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
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
