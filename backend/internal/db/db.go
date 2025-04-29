package db

import (
	"database/sql"
	_"modernc.org/sqlite"
)

func SetupDb(filePath string) (*sql.DB, error){
	db, err := sql.Open("sqlite", "file:///home/jrport/repos/chat/backend/" + filePath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		print("Error on database ping.")
		return nil, err
	}

	return db, nil
}
