package registration

import (
	"chat/api/database/models"
	"chat/api/handles/utils"
	handlesUtils "chat/api/handles/utils"
	"chat/api/services"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegistrationHandle(w http.ResponseWriter, r *http.Request) *utils.ResponseError {
	if r.Method != http.MethodPost{
		return handlesUtils.NewResponseError("Invalid method", http.StatusMethodNotAllowed) 
	}

	if ct := r.Header.Get("Content-Type"); ct != "application/json"{
		return handlesUtils.NewResponseError("Invalid Content-Type", http.StatusBadRequest)
	}

	formBuffer, err := io.ReadAll(r.Body)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusBadRequest)
	}
	userRegistration, err := handlesUtils.SerializeRegistration(&formBuffer)

	log.Printf("d: %v", userRegistration)
	userInDb, err := models.GetUser(userRegistration)
	if err != nil && err != mongo.ErrNoDocuments {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	if err == nil{
		switch {
		case userInDb.Confirmed:
			return handlesUtils.NewResponseError(NewUserRegistrationError(EmailAlreadyInUse).Msg, http.StatusConflict)
		case !userInDb.Confirmed && time.Since(userInDb.UpdatedAt) < time.Duration(time.Minute * 2):
			return handlesUtils.NewResponseError(NewUserRegistrationError(TokenIssuedRecently).Msg, http.StatusTooManyRequests)
		default:
			err := services.IssueToken(&userInDb.ID)
			if err != nil {
				return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
			}
			return nil
		}
	}

	log.Printf("c: %v", userRegistration)
	user, err := models.CreateUser(userRegistration)
	err = services.IssueToken(user)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}

type RegistrationError int

const (
	TokenIssuedRecently RegistrationError = iota
	EmailAlreadyInUse
)

type UserRegistrationError struct {
	Msg string
}

func NewUserRegistrationError(kind RegistrationError) *UserRegistrationError{
	return nil
}

func (u *UserRegistrationError)Error() string{
	return u.Msg
}
