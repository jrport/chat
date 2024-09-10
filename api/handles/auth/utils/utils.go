package utils

import (
	"encoding/json"
	"log/slog"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistration struct {
	ID                   primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Email                string             `json:"email" bson:"Email"`
	Password             string             `json:"password" bson:"Password"`
	PasswordConfirmation string             `json:"password-confirmation" bson:"-"`
	Confirmed            bool               `json:"-" bson:"Confirmed"`
	CreatedAt            time.Time          `json:"-" bson:"CreatedAt"`
	UpdatedAt            time.Time          `json:"-" bson:"UpdatedAt"`
}

func (u *UserRegistration)Hash() error{
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    u.Password = string(hashedPassword)
    return nil
}

func SerializeRegistration(rawJson *[]byte) (*UserRegistration, error) {
	var decodedUserRegistration UserRegistration
	err := json.Unmarshal(*rawJson, &decodedUserRegistration)
	if err != nil {
		return nil, err
	}

	
	if valError := decodedUserRegistration.Validate(); valError != nil {
		return nil, valError
	}

	decodedUserRegistration.Confirmed = false
	decodedUserRegistration.CreatedAt = time.Now()
	decodedUserRegistration.UpdatedAt = time.Now()

	return &decodedUserRegistration, nil
}

type Criteria int

const (
	EmptyEmail          Criteria = iota
	EmptyPassword       Criteria = iota
	NonMatchingPassword Criteria = iota
)

type ValidationError struct {
	Kind Criteria
}

func NewValidationError(kind Criteria) *ValidationError {
	return &ValidationError{
		Kind: kind,
	}
}

func (v ValidationError) Error() string {
	switch v.Kind {
	case EmptyEmail:
		return "Empty Email"
	case EmptyPassword:
		return "Empty Password"
	case NonMatchingPassword:
		return "Password and Password confirmation must Match"
	default:
		return "Unknown error on registration validation"
	}
}

func (u *UserRegistration) Validate() *ValidationError {
	switch {
	case IsEmpty(u.Email):
		return NewValidationError(EmptyEmail)
	case IsEmpty(u.Password):
		return NewValidationError(EmptyPassword)
	case u.Password == u.PasswordConfirmation:
		return NewValidationError(NonMatchingPassword)
	default:
		return nil
	}

}

func IsEmpty(field string) bool {
	return strings.TrimSpace(field) == ""
}

func CheckCredentials(inputLogin, storedCredentials *UserRegistration) bool{
	if err := bcrypt.CompareHashAndPassword([]byte(storedCredentials.Password), []byte(inputLogin.Password)); err != nil {
		slog.Error("Error on login check " + err.Error())
		return false
	}
	return true
}
