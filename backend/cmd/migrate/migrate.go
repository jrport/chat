package main

import (
	"fmt"
	"jport/chat/backend/internal/db"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

const (
	migrationsPath = "file:///home/jrport/repos/chat/backend/migrations/"
	databaseName   = "chat"
)

func runMigrations(m *migrate.Migrate) error {
	println("Migrating...")
	if err := m.Steps(1); err != nil {
		if err == migrate.ErrNoChange {
			return fmt.Errorf("No migrations applied: %v", err.Error())
		}
		return fmt.Errorf("Error running migration!\n%v\n", err.Error())
	}
	println("Sucessfully applied 1 migration!")
	return nil
}

func fixDirtyVersion(desiredVersion int, m *migrate.Migrate) error {
	if err := m.Force(desiredVersion); err != nil {
		return fmt.Errorf("Error fixing dirty db version: %v", err.Error())
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		panic(` Please provide the name of
			the sqlite database file to 
			seeded.`)
	}

	dbConn, err := db.SetupDb(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err = dbConn.Close(); err != nil {
			panic(fmt.Sprintf("Erro closing connection...\nError -> %v\n", err.Error()))
		}
	}()

	driver, err := sqlite.WithInstance(dbConn, &sqlite.Config{})
	if err != nil {
		panic("Erro on migration setup!")
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		panic(fmt.Sprintf("Error instantiating migration: %v", err.Error()))
	}

	switch {
	case len(os.Args) == 2:
		runMigrations(m)

	case len(os.Args) == 3:
		desiredVersion, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(fmt.Sprintf("Invalid version to force: %v", err.Error()))
		}

		if err = fixDirtyVersion(desiredVersion, m); err != nil {
			panic(err.Error())
		}
	}

}
