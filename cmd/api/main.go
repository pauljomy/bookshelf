package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pauljomy/bookshelf/internal/env"
)

const version = "1.0.0"

type application struct {
	logger *slog.Logger
	cfg    *config
}

type config struct {
	port int
	env  string
	db   dbConfig
}

type dbConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load .env file", "error", err)
		os.Exit(1)
	}

	var cfg config

	port := env.GetInt("PORT", 4040)
	environment := env.GetString("ENV", "development")
	dsn := env.GetString("DSN", "")
	maxOpenConns := env.GetInt("MAX_OPEN_CONNS", 25)
	maxIdleConns := env.GetInt("MAX_IDLE_CONNS", 25)
	maxIdleTime := env.GetString("MAX_IDLE_TIME", "5m")

	flag.IntVar(&cfg.port, "port", port, "port to listen on")
	flag.StringVar(&cfg.env, "env", environment, "environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", dsn, "database connection string")
	flag.IntVar(&cfg.db.maxOpenConns, "max-open-conns", maxOpenConns, "maximum number of open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "max-idle-conns", maxIdleConns, "maximum number of idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "max-idle-time", maxIdleTime, "maximum idle time")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg.db.dsn, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Error("Failed to open database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
		cfg:    &cfg,
	}

	logger.Info("Server starting", "at port", cfg.port)
	logger.Info("Database connection established")

	err = app.server()
	if err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}

}

func openDB(dsn string, maxOpenConns int, maxIdleConns int, maxIdleTime string) (*pgxpool.Pool, error) {
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(maxOpenConns)
	poolConfig.MinConns = int32(maxIdleConns)
	poolConfig.MaxConnIdleTime = duration

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
