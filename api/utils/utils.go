package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email        string    `json:"email" bson:"email"`
	PasswordHash string    `json:"password" bson:"password_hash"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
}


func (c *Credentials) Hash() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(c.PasswordHash), bcrypt.DefaultCost)
	c.PasswordHash = string(hashed)
}

func ReadJson(r *io.ReadCloser) (*Credentials, error) {
	credentials := new(Credentials)
	decoder := json.NewDecoder(*r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&credentials)
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func WriteJson(w http.ResponseWriter, credentials *Credentials) error {
	err := json.NewEncoder(w).Encode(credentials)
	if err != nil {
		return err
	}
	return nil
}
