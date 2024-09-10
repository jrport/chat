package registration

import (
	"chat/api/database/models"
	"chat/api/handles/utils"
	handlesUtils "chat/api/handles/utils"
	authUtils "chat/api/handles/auth/utils"
	"chat/api/services/mailer"
	"chat/api/services/token"
	tokenService "chat/api/services/token"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)


func RegistrationHandleFunc(w http.ResponseWriter, r *http.Request, m *mailer.Mailer) *utils.ResponseError {
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
	userRegistration, err := authUtils.SerializeRegistration(&formBuffer)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

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
			_, err := token.IssueToken(&userInDb.ID, token.Confirmation)
			if err != nil {
				return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
			}
		}
	}

	user, err := models.CreateUser(userRegistration)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}
	verificationToken, err := token.IssueToken(user, token.Confirmation)
	if err != nil {
		return handlesUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	mailOrder := mailer.NewMailerOrder(
		userRegistration.Email,
		mailer.ValidationTokenMail,
		&map[string]string{
			"verificationToken": *verificationToken,
		},
	)
	m.IssueMail(*mailOrder)

    fmt.Fprint(w, "User created successfully, please verify your email to proceed!\n")
	return nil
}

func ValidationHandle(w http.ResponseWriter, r *http.Request) *utils.ResponseError {
    if r.Method != http.MethodGet {
        return utils.NewResponseError("Invalid Method", http.StatusMethodNotAllowed)
    }
    
    token := r.URL.Query().Get("token")
    
    if (token == "") {
        return utils.NewResponseError("Invalid parameters", http.StatusBadRequest)
    }
    
    storedId, err := tokenService.ValidateToken(token)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }

    err = models.ValidateUser(*storedId)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }

    err = tokenService.ExpireToken(token)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Vc foi viadado com sucesso!")
    return nil
}

func PasswordResetHandle(w http.ResponseWriter, r *http.Request) *utils.ResponseError{
	return nil
}
