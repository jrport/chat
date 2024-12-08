package handlers

import (
	"fmt"
	"jport/chat/backend/internal/errors"
	"jport/chat/backend/internal/json"
	"net/http"
	"strings"
)

func loginHandler(_ http.ResponseWriter, r *http.Request) error {
	credentials, err := json.GetUserLogin(r.Body)
	if err != nil {
		return err
	}

	if strings.TrimSpace(credentials.Login) == "" || strings.TrimSpace(credentials.Password) == "" {
		return errors.NewHttpError(http.StatusUnauthorized, "No empty fields.")
	}

	print(fmt.Sprintf("Login: %v\nPassword: %v\n", credentials.Login, credentials.Password))
	return nil
}

func registrationHandler(_ http.ResponseWriter, r *http.Request) error {
	credentials, err := json.GetUserRegistration(r.Body)
	if err != nil {
		return err
	}

	if strings.TrimSpace(credentials.Email) == "" || strings.TrimSpace(credentials.Password) == "" || strings.TrimSpace(credentials.Login) == "" || strings.TrimSpace(credentials.Login) == "" {
		return errors.NewHttpError(http.StatusUnauthorized, "No empty fields.")
	}

	if strings.TrimSpace(credentials.Login) == "" || strings.TrimSpace(credentials.Password) == "" {
		return errors.NewHttpError(http.StatusBadRequest, "Password and Password Confirmation must be equal.")
	}

	// if err := users.CreateUser(credentials.Email, credentials.Login, credentials.Password); err != nil {
	// 	return err
	// }

	return nil
}
