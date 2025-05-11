package main

import (
	"chat/backend/internal/db"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getDbConnection() *migrate.Migrate{
    db, err := sql.Open("sqlite", db.DbURL)
    if err != nil {
        fmt.Printf("Error opening db file: %v", err.Error())
        return nil
    }
    driver, err := sqlite.WithInstance(db, &sqlite.Config{})
    if err != nil {
        fmt.Printf("Error instatiating sqlite driver: %v", err.Error())
        return nil
    }
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres", driver)
    if err != nil {
        fmt.Printf("Error migration folder: %v", err.Error())
        return nil
    }
    return m
}

func main() {
	cmd := os.Args[1]
    switch cmd {
        case "up":
            m := getDbConnection()
            if err := m.Up(); err != nil {
                fmt.Printf("Error on up: %v", err.Error())
            }
            return
        case "down":
            m := getDbConnection()
            if err := m.Steps(-1); err != nil {
                fmt.Printf("Error on rollback: %v", err.Error())
            }
            return
        case "force":
            if (len(os.Args) <= 2) {
                println("Not enough args for FORCE stamp op. Please input version.")
                return
            }
            version, _ := strconv.Atoi(os.Args[2])
            m := getDbConnection()
            if err := m.Force(version); err != nil {
                fmt.Printf("Error on stamping: %v", err.Error())
            }
            return
        default:
            println("Unknown cmd.")
            return
    }
}
