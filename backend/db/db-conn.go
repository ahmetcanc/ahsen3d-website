package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// ConnectDB establishes a PostgreSQL connection using pgxpool and runs initial SQL file
func ConnectDB() error {
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return fmt.Errorf("error parsing database config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	// Store globally
	DB = pool

	fmt.Println("✅ Database connection established successfully")

	// Run SQL initialization
	if err := runInitSQL(ctx, pool); err != nil {
		return fmt.Errorf("error running init SQL file: %w", err)
	}

	return nil
}

func runInitSQL(ctx context.Context, pool *pgxpool.Pool) error {
	sqlFile := "db/database-init.sql"

	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %w", sqlFile, err)
	}

	sql := string(sqlBytes)

	_, err = pool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("failed to execute init SQL: %w", err)
	}

	fmt.Println("✅ database-init.sql executed successfully")
	return nil
}
