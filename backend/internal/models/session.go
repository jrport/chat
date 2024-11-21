package models

import "database/sql"

func GetSession(db *sql.DB) error {
	db.QueryRowContext("SELECT '")
}
