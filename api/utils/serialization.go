package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ReadJson(r *io.ReadCloser) (*Credentials, error) {
	credentials := new(Credentials)
	decoder := json.NewDecoder(*r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&credentials)
	if err != nil {
		return nil, err
	}
	switch {
	case credentials.Email == "":
		return nil, fmt.Errorf("Email é obrigatório")
	case credentials.PasswordHash == "":
		return nil, fmt.Errorf("Senha é obrigatória")
	default:
        credentials.Confirmed = false
		return credentials, nil
	}
}

func WriteJson(w http.ResponseWriter, credentials *Credentials) error {
	err := json.NewEncoder(w).Encode(credentials)
	if err != nil {
		return err
	}
	return nil
}
