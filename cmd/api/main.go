package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/mystpen/test-task-Mobydev/config"
	"github.com/mystpen/test-task-Mobydev/internal/logger"
)

type application struct {
	config config.Config
	logger *logger.Logger
}

func main() {
	logger := logger.NewLogger()

	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	defer db.Close()

	logger.InfoLog.Printf("database connection pool established")

	// Database migrations
	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("./migrations", "postgres", migrationDriver)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.ErrLog.Fatal(err)
	}

	logger.InfoLog.Printf("database migrations applied")

	app := &application{
		config: cfg,
		logger: &logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.InfoLog.Printf("starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	logger.ErrLog.Fatal(err)
}
