package db

import (
	"database/sql"
	"fmt"
	"jport/chat/backend/internal/errors"
	"jport/chat/backend/internal/hash"
	"net/http"
)

func CreateUser(db *sql.DB, username, email, password string) (*int64, error) {
	count := 0

	err := db.QueryRow(`
		SELECT COUNT(p.*) 
		FROM users 
		WHERE username = :username`,
		sql.Named("username", username)).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if count != 0 {
		return nil, errors.NewHttpError(http.StatusConflict, "Username already in use.")
	}

	db.QueryRow("SELECT COUNT(p.*) FROM users WHERE email = :email", sql.Named("email", email)).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if count != 0 {
		return nil, errors.NewHttpError(http.StatusConflict, "Email already in use.")
	}

	hashedPassword := hash.HashPassword(password)
	res, err := db.Exec(
		`INSERT INTO users (username, password, email)
		VALUES (:username, :password, :email)`,
		sql.Named("username", username),
		sql.Named("password", hashedPassword),
		sql.Named("email", email),
	)
	if err != nil {
		return nil, errors.NewHttpError(
			http.StatusInternalServerError,
			fmt.Sprintf(
				"Error on insertion of record: login=%v password=%v email=%v",
				username, password, email,
			),
		)
	}
	id, _ := res.LastInsertId()

	return &id, nil
}
