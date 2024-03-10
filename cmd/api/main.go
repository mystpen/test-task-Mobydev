package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/mystpen/test-task-Mobydev/config"
	"github.com/mystpen/test-task-Mobydev/internal/logger"
	"github.com/mystpen/test-task-Mobydev/internal/repository"
	"github.com/mystpen/test-task-Mobydev/internal/rest/handler"
	"github.com/mystpen/test-task-Mobydev/internal/service"
)

func main() {
	logger := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil{
		logger.ErrLog.Fatal(err)
	}

	// Connect to DB
	db, err := openDB(*cfg)
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
	migrator, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", migrationDriver)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.ErrLog.Fatal(err)
	}

	// defer migrator.Down()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, &logger)

	logger.InfoLog.Printf("database migrations applied")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.InfoLog.Printf("starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	logger.ErrLog.Fatal(err)
}
