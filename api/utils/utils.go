package utils

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/goloop/env"
	"golang.org/x/crypto/bcrypt"
)

type MailCredentials struct {
	Email    string
	AppToken string
	Host     string
}

func GetMailerCredentials() (*MailCredentials, error) {
	if err := env.Load(".env"); err != nil {
		return nil, err
	}

	return &MailCredentials{
		Email:    env.Get("email"),
		AppToken: env.Get("password"),
		Host:     env.Get("host"),
	}, nil
}

type DecodingError struct{}

func (d DecodingError) Error() string {
	return "Error on parsing json: Fields cannot be empty"
}

type Account struct {
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"password" bson:"passwordHash"`
	Confirmed    bool   `json:"-" bson:"confirmed"`
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
