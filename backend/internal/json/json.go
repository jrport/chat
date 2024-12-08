package json

import (
	"encoding/json"
	"errors"
	customError"jport/chat/backend/internal/errors"
	"io"
)

type UserLogin struct {
	Login string 
	Password string 
}

type UserRegistration struct {
	Login string
	Email string
	Password string
	PasswordConfirmation string
}


func GetUserLogin(body io.ReadCloser) (*UserLogin, error){
	credentials := &UserLogin{}
	err := json.NewDecoder(body).Decode(credentials)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, customError.AuthenticationError{}
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
			return nil, customError.AuthenticationError{}
		}
		return nil, err
	}

	if credentials.Password != credentials.PasswordConfirmation {
		return nil, customError.PasswordConfirmationMismatchError{}
	}
	return credentials, nil
}



//
// type UserPasswordRecovery struct {
// 	Login string
// 	Email string
// }
//
