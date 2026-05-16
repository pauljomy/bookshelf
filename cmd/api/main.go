package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/pauljomy/bookshelf/internal/env"
)

type application struct {
	logger *slog.Logger
	cfg    *config
}

type config struct {
	port int
}

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load .env file", "error", err)
		os.Exit(1)
	}

	var cfg config

	port := env.GetInt("PORT", 4040)

	flag.IntVar(&cfg.port, "port", port, "port to listen on")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
		cfg:    &cfg,
	}

	logger.Info("Server starting", "at port", cfg.port)

	err = app.server()
	if err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}

}
