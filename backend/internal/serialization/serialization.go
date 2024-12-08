package serialization

import (
	"encoding/json"
	"errors"
	"io"
	customError "jport/chat/backend/internal/errors"
	"net/http"
)

type UserLogin struct {
	Username string 
	Password string 
}

type UserRegistration struct {
	Username string
	Email string
	Password string
	PasswordConfirmation string
}


func GetUserLogin(body io.ReadCloser) (*UserLogin, error){
	credentials := &UserLogin{}
	err := json.NewDecoder(body).Decode(credentials)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, customError.NewHttpError(http.StatusBadRequest, "All fields are obligatory.")
		}
		return nil, err
	}
	return credentials, nil
}

func GetUserRegistration(body io.ReadCloser) (*UserRegistration, error){
	credentials := &UserRegistration{}
	err := json.NewDecoder(body).Decode(credentials)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, customError.NewHttpError(http.StatusBadRequest, "All fields are obligatory.")
		}
		return nil, err
	}

	if credentials.Password != credentials.PasswordConfirmation {
		return nil, customError.NewHttpError(http.StatusBadRequest, "Password and Password Confirmation must be match.")
	}

	return credentials, nil
}

