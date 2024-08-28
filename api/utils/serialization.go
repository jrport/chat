package utils

import (
	"encoding/json"
	"io"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type DecodingError struct{}

func (d DecodingError) Error() string {
	return "Error on parsing json: Fields cannot be empty"
}

type Account struct {
    ID           primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"password" bson:"passwordHash"`
	Confirmed    bool               `json:"-" bson:"confirmed"`
}

func (u *Account) Hash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func JsonToCredentials(requestBody *io.ReadCloser) (*Account, error) {
	credentials := Account{Confirmed: false}

	decoder := json.NewDecoder(*requestBody)
	err := decoder.Decode(&credentials)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(credentials.Email) == "" || strings.TrimSpace(credentials.PasswordHash) == "" {
		return nil, DecodingError{}
	}

	err = credentials.Hash()
	if err != nil {
		return nil, err
	}

	return &credentials, nil
}
