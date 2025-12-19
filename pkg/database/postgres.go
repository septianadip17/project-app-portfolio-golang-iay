package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresConnection membuat koneksi ke PostgreSQL menggunakan pgxpool
func NewPostgresConnection(connString string) (*pgxpool.Pool, error) {
	// Konfigurasi koneksi
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("gagal parsing config database: %w", err)
	}

	// Setting timeout dsb (optional)
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	// Buka koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("gagal connect ke database: %w", err)
	}

	// Tes Ping
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	return pool, nil
}
