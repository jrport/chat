package sessions

import (
	"chat/api/database/models"
	authUtils "chat/api/handles/auth/utils"
	handleUtils "chat/api/handles/utils"
	tokenService "chat/api/services/token"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewSessionHandle(w http.ResponseWriter, r *http.Request) *handleUtils.ResponseError {
	if r.Method != http.MethodPost {
		return handleUtils.NewResponseError("Invalid method", http.StatusMethodNotAllowed)
	}

	if ct := r.Header.Get("Content-Type"); ct != "application/json" {
		return handleUtils.NewResponseError("Invalid Content-Type", http.StatusBadRequest)
	}

	formBuffer, err := io.ReadAll(r.Body)
	if err != nil {
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	userLogin, err := authUtils.SerializeRegistration(&formBuffer)
	if err != nil {
		switch err.(type) {
		case authUtils.ValidationError:
			return handleUtils.NewResponseError(err.Error(), http.StatusBadRequest)
		default:
			return handleUtils.NewResponseError(err.Error(), http.StatusBadRequest)
		}
	}

	userInDb, err := models.GetUser(userLogin)

	switch {
	case err != nil && err != mongo.ErrNoDocuments:
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	case err != nil:
		return handleUtils.NewResponseError("Wrong Credentials", http.StatusUnauthorized)
	case !userInDb.Confirmed:
		return handleUtils.NewResponseError("Please confirm your email", http.StatusForbidden)
	case !authUtils.CheckCredentials(userLogin, userInDb):
		return handleUtils.NewResponseError("Wrong Credentials", http.StatusUnauthorized)
	}

	sessToken, err := tokenService.IssueToken(&userInDb.ID, tokenService.Session)
	if err != nil {
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     "sessionId",
		Value:    *sessToken,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: false,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfull login! Welcome.")
	slog.Info("Sucessfull login from user " + userInDb.Email)

	return nil
}

func TerminateSessionHandle(w http.ResponseWriter, r *http.Request) *handleUtils.ResponseError {
	if r.Method != http.MethodDelete {
		return handleUtils.NewResponseError("Invalid Method", http.StatusMethodNotAllowed)
	}

	// token := r.URL.Query().Get("token")
	sessionCookie, err := r.Cookie("sessionId")
	if err != nil {
		return handleUtils.NewResponseError("Invalid parameters", http.StatusBadRequest)
	}

	sessionToken := sessionCookie.Value
	_, err = tokenService.ValidateToken(sessionToken)
	if err != nil {
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	err = tokenService.ExpireToken(sessionToken)
	if err != nil {
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	slog.Info("Session terminated sucessfully")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Session terminated sucessfully")
	return nil
}

func ValidateSession(w http.ResponseWriter, r *http.Request) *handleUtils.ResponseError {
	if r.Method != http.MethodGet {
		return handleUtils.NewResponseError("Invalid Method", http.StatusMethodNotAllowed)
	}

	token := r.URL.Query().Get("token")

	if token == "" {
		return handleUtils.NewResponseError("Invalid parameters", http.StatusBadRequest)
	}

	_, err := tokenService.ValidateToken(token)
	if err != nil {
		return handleUtils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	return nil
}
