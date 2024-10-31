package main

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// Embed the migrations
//
//go:embed sql/schema/*.sql
var migrations embed.FS

func applyMigrations(db *sql.DB) error {
	// Set the base file system for Goose to embedded FS
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	// Run migrations
	if err := goose.Up(db, "sql/schema"); err != nil {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	return nil
}
