package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"samurenkoroma/services/internal/infrastructure/configs"
	"samurenkoroma/services/pkg/db"
	"sort"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var (
	moduleFlag = flag.String("module", "all", "Module to migrate (shared, farm, crop, growing, all)")
	direction  = flag.String("direction", "up", "Migration direction (up, down)")
	steps      = flag.Int("steps", 0, "Number of steps for up/down (0 = all)")
	dsn        = flag.String("dsn", "", "PostgreSQL DSN (or use DATABASE_URL env)")
	force      = flag.Int64("force", 0, "Force set version")
	verbose    = flag.Bool("verbose", false, "Verbose output")
)

func main() {
	flag.Parse()

	conf := configs.LoadConfig()
	ctx := context.Background()

	conn, err := db.NewDB(ctx, conf.Db.Dsn)
	if err != nil {
		return
	}
	defer conn.Close()

	// Создаём таблицу миграций
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		id         SERIAL PRIMARY KEY,
		module     TEXT NOT NULL,
		version    BIGINT NOT NULL,
		applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		dirty      BOOLEAN NOT NULL DEFAULT FALSE,
		UNIQUE(module, version)
	)`
	if _, err := conn.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create migrations table: %v", err)
	}

	modules := getModules(*moduleFlag)

	for _, mod := range modules {
		if *force > 0 {
			if err := forceVersion(conn, mod.Name, *force); err != nil {
				log.Fatalf("Failed to force version for %s: %v", mod.Name, err)
			}
			continue
		}

		migrationsPath := filepath.Join("internal", "modules", mod.Path)
		if err := runMigrations(conn, mod.Name, migrationsPath, *direction, *steps, *verbose); err != nil {
			log.Fatalf("Failed to run migrations for %s: %v", mod.Name, err)
		}
	}

	log.Println("All operations completed successfully!")
}

func runMigrations(db *sql.DB, moduleName, migrationsPath, direction string, steps int, verbose bool) error {
	// Проверяем существование папки
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		if verbose {
			log.Printf("Directory %s does not exist, skipping", migrationsPath)
		}
		return nil
	}

	// Сканируем файлы миграций
	files, err := filepath.Glob(filepath.Join(migrationsPath, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("failed to list migration files: %w", err)
	}

	if len(files) == 0 {
		if verbose {
			log.Printf("No migration files found for %s", moduleName)
		}
		return nil
	}

	// Парсим миграции
	migrations := make([]Migration, 0)
	for _, f := range files {
		base := filepath.Base(f)
		parts := strings.SplitN(base, "_", 2)
		if len(parts) != 2 {
			continue
		}
		version, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}
		upPath := f
		downPath := filepath.Join(migrationsPath, strings.TrimSuffix(base, ".up.sql")+".down.sql")
		migrations = append(migrations, Migration{
			Version:  version,
			UpPath:   upPath,
			DownPath: downPath,
		})
	}

	// Сортируем по версии
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	// Получаем все применённые версии для этого модуля
	appliedVersions, err := getAppliedVersions(db, moduleName)
	if err != nil {
		return fmt.Errorf("failed to get applied versions: %w", err)
	}

	if verbose {
		log.Printf("Applied versions for %s: %v", moduleName, appliedVersions)
	}

	switch direction {
	case "up":
		return upMigrations(db, moduleName, migrations, appliedVersions, steps, verbose)
	case "down":
		return downMigrations(db, moduleName, migrations, appliedVersions, steps, verbose)
	default:
		return fmt.Errorf("unknown direction: %s", direction)
	}
}

func upMigrations(db *sql.DB, moduleName string, migrations []Migration, appliedVersions map[int64]bool, steps int, verbose bool) error {
	// Находим неприменённые миграции
	var toApply []Migration
	for _, m := range migrations {
		if !appliedVersions[m.Version] {
			toApply = append(toApply, m)
		}
	}

	if steps > 0 && steps < len(toApply) {
		toApply = toApply[:steps]
	}

	if len(toApply) == 0 {
		if verbose {
			log.Printf("No pending migrations for %s", moduleName)
		}
		return nil
	}

	for _, m := range toApply {
		if verbose {
			log.Printf("Applying migration %d for %s", m.Version, moduleName)
		}

		// Читаем SQL
		content, err := os.ReadFile(m.UpPath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", m.UpPath, err)
		}

		// Выполняем в транзакции
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute migration %d: %w", m.Version, err)
		}

		// Записываем в таблицу миграций
		if err := insertMigration(tx, moduleName, m.Version); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration: %w", err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}

		log.Printf("✓ Applied migration %d for %s", m.Version, moduleName)
	}

	return nil
}

func downMigrations(db *sql.DB, moduleName string, migrations []Migration, appliedVersions map[int64]bool, steps int, verbose bool) error {
	// Находим применённые миграции (в обратном порядке)
	var toRollback []Migration
	for i := len(migrations) - 1; i >= 0; i-- {
		if appliedVersions[migrations[i].Version] {
			toRollback = append(toRollback, migrations[i])
		}
	}

	if steps > 0 && steps < len(toRollback) {
		toRollback = toRollback[:steps]
	}

	if len(toRollback) == 0 {
		if verbose {
			log.Printf("No migrations to rollback for %s", moduleName)
		}
		return nil
	}

	for _, m := range toRollback {
		if verbose {
			log.Printf("Rolling back migration %d for %s", m.Version, moduleName)
		}

		// Проверяем наличие down-файла
		if _, err := os.Stat(m.DownPath); os.IsNotExist(err) {
			return fmt.Errorf("down migration file missing for version %d: %s", m.Version, m.DownPath)
		}

		content, err := os.ReadFile(m.DownPath)
		if err != nil {
			return fmt.Errorf("failed to read down migration file: %w", err)
		}

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to execute down migration %d: %w", m.Version, err)
		}

		// Удаляем запись из таблицы миграций
		if err := deleteMigration(tx, moduleName, m.Version); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete migration record: %w", err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}

		log.Printf("✓ Rolled back migration %d for %s", m.Version, moduleName)
	}

	return nil
}

func getAppliedVersions(db *sql.DB, moduleName string) (map[int64]bool, error) {
	rows, err := db.Query(`
		SELECT version FROM schema_migrations
		WHERE module = $1
	`, moduleName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	versions := make(map[int64]bool)
	for rows.Next() {
		var version int64
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		versions[version] = true
	}
	return versions, nil
}

func insertMigration(tx *sql.Tx, moduleName string, version int64) error {
	_, err := tx.Exec(`
		INSERT INTO schema_migrations (module, version)
		VALUES ($1, $2)
	`, moduleName, version)
	return err
}

func deleteMigration(tx *sql.Tx, moduleName string, version int64) error {
	_, err := tx.Exec(`
		DELETE FROM schema_migrations
		WHERE module = $1 AND version = $2
	`, moduleName, version)
	return err
}

func forceVersion(db *sql.DB, moduleName string, version int64) error {
	// Удаляем все записи для модуля
	if _, err := db.Exec(`DELETE FROM schema_migrations WHERE module = $1`, moduleName); err != nil {
		return err
	}
	if version > 0 {
		// Добавляем указанную версию
		if _, err := db.Exec(`
			INSERT INTO schema_migrations (module, version)
			VALUES ($1, $2)
		`, moduleName, version); err != nil {
			return err
		}
		log.Printf("Forced version %d for %s", version, moduleName)
	} else {
		log.Printf("Cleared all migrations for %s", moduleName)
	}
	return nil
}

type Migration struct {
	Version  int64
	UpPath   string
	DownPath string
}

type module struct {
	Name string
	Path string
}

func getModules(moduleFlag string) []module {
	allModules := []module{
		//{Name: "shared", Path: "shared"},
		{Name: "farm", Path: "farm/infrastructure/persistence/postgres/migrations"},
		{Name: "crop", Path: "crop/infrastructure/persistence/postgres/migrations"},
		{Name: "growing", Path: "growing/infrastructure/persistence/postgres/migrations"},
	}

	if moduleFlag == "all" {
		return allModules
	}

	for _, m := range allModules {
		if m.Name == moduleFlag {
			return []module{m}
		}
	}

	log.Fatalf("Unknown module: %s. Available: shared, farm, crop, growing, all", moduleFlag)
	return nil
}
