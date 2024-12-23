package main

import (
	"flag"
	"log"

	"github.com/VadimShara/coursework_DB_shop/internal/config"
	"github.com/VadimShara/coursework_DB_shop/internal/repo"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	runMigrations()
}

func runMigrations() {
	PostgresCfg := config.Load().PGConfig
	dbUrl := repo.PGConnectionStr(PostgresCfg)

	// Parse command-line flags
	migrateCmd := flag.String("migrate", "", "Specify 'up' or 'down' to run migrations")
	flag.Parse()

	// Run migrations
	m, err := migrate.New(
		"file://migrations/",
		dbUrl,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Perform migration based on command-line flag
	switch *migrateCmd {
	case "up":
		migrateUp(m)
	case "down":
		migrateDown(m)
	default:
		log.Fatal("--migrate flag not specified. Use '--migrate=up' or '--migrate=down'")
	}

	log.Println("Migrations applied successfully")
}

func migrateUp(m *migrate.Migrate) {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}

func migrateDown(m *migrate.Migrate) {
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
}