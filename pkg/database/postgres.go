package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(connString string) (*pgxpool.Pool, error) {
	// Konfigurasi koneksi
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("gagal parse config DB: %w", err)
	}

	// Setting timeout koneksi
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	// Coba koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat connection pool: %w", err)
	}

	// Tes Ping ke DB
	if err := dbPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	return dbPool, nil
}
