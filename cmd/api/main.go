package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mystpen/test-task-Mobydev/config"
)


type application struct {
	config config.Config
	logger *log.Logger
}

func main() {
	var cfg config.Config

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
