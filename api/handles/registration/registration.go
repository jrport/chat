package registration

import (
	"chat/api/database/models"
	"chat/api/handles/utils"
	handlesUtils "chat/api/handles/utils"
	"chat/api/services/token"
	"fmt"
	"io"
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

	userInDb, err := models.GetUser(userRegistration)
	if err != nil && err != mongo.ErrNoDocuments {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	if err == nil{
		switch {
		case userInDb.Confirmed:
            err := NewUserRegistrationError(EmailAlreadyInUse)
			return handlesUtils.NewResponseError(err.Error(), http.StatusConflict)
		case !userInDb.Confirmed && time.Since(userInDb.UpdatedAt) < time.Duration(time.Minute * 2):
            error := NewUserRegistrationError(TokenIssuedRecently)
			return handlesUtils.NewResponseError(error.Error(), http.StatusTooManyRequests)
		default:
			err := token.IssueToken(&userInDb.ID)
			if err != nil {
				return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
			}
			return nil
		}
	}

	user, err := models.CreateUser(userRegistration)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}
	err = token.IssueToken(user)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

    fmt.Fprint(w, "User created successfully, please verify your email to proceed")
	return nil
}
