package handlers

import (
	"database/sql"
	"jport/chat/backend/internal/errors"
	"jport/chat/backend/internal/serialization"
	"net/http"
	"strings"
)

func LoginHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	credentials, err := serialization.GetUserLogin(r.Body)
	if err != nil {
		return err
	}

	if strings.TrimSpace(credentials.Username) == "" {
		return errors.NewHttpError(http.StatusUnauthorized, "Username is obligatory.")
	}
	if strings.TrimSpace(credentials.Password) == "" {
		return errors.NewHttpError(http.StatusUnauthorized, "Password is obligatory.")
	}

	return nil
}

func RegistrationHandler(db *sql.DB, _ http.ResponseWriter, r *http.Request) error {
	credentials, err := serialization.GetUserRegistration(r.Body)
	if err != nil {
		return err
	}

	if strings.TrimSpace(credentials.Email) == "" {
		return errors.NewHttpError(http.StatusBadRequest, "Email can't be empty")
	}

	if strings.TrimSpace(credentials.Username) == "" {
		return errors.NewHttpError(http.StatusBadRequest, "Usename can't be empty")
	}

	if strings.TrimSpace(credentials.Password) == "" {
		return errors.NewHttpError(http.StatusBadRequest, "Password can't be empty")
	}

	if strings.TrimSpace(credentials.Username) == "" || strings.TrimSpace(credentials.Password) == "" {
		return errors.NewHttpError(http.StatusBadRequest, "Password and Password Confirmation must be equal.")
	}

	// TODO AQUi
	// if err := users.CreateUser(credentials.Email, credentials.Login, credentials.Password); err != nil {
	// 	return err
	// }
	return nil
}
