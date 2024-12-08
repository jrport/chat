package db

import (
	"database/sql"
	"jport/chat/backend/internal/errors"
	"net/http"
)

func CreateUser(db *sql.DB, username, email, password string) error {
	count := 0

	err := db.QueryRow(`
		SELECT COUNT(p.*) 
		FROM users 
		WHERE username = :username`,
		sql.Named("username", username)).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if count != 0 {
		return errors.NewHttpError(http.StatusConflict, "Username already in use.")
	}

	db.QueryRow("SELECT COUNT(p.*) FROM users WHERE email = :email", sql.Named("email", email)).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if count != 0 {
		return errors.NewHttpError(http.StatusConflict, "Email already in use.")
	}

	// TODO GETTING ENV VARS, HASHING PASSWORD AND STORING EVERYTHING
	return nil
}
