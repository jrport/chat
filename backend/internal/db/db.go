package db

import (
	"database/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func SetupDb(filePath string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", "file:" + filePath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
