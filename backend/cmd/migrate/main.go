package main

import (
	"database/sql"
	"os"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if argc := len(os.Args); argc < 2 { 
		log.Fatal("Specify up or down migration")	
	}

	command := os.Args[1]

	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("db file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	switch command{
		case "up":
			if err := m.Up(); err != nil {
				log.Fatal(err)
			}
			log.Print("Finished migration")
		case "down":
			if err := m.Down(); err != nil {
				log.Fatal(err)
			}
			log.Print("Finished migration")
	}
}
