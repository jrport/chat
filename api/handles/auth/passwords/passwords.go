package passwords

import (
	"chat/api/database/models"
	authUtils "chat/api/handles/auth/utils"
	"chat/api/handles/utils"
	"chat/api/services/mailer"
	"chat/api/services/token"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func PasswordResetTokenHandle(w http.ResponseWriter, r *http.Request, m *mailer.Mailer) *utils.ResponseError {
	if r.Method != http.MethodPost {
		return utils.NewResponseError("Invalid method", http.StatusMethodNotAllowed)
	}

	ct := r.Header.Get("Content-Type")
	if ct == "" {
		return utils.NewResponseError("Bad Content-Type", http.StatusBadRequest)
	}

	formBuffer, err := io.ReadAll(r.Body)
	if err != nil {
		return utils.NewResponseError(err.Error(), http.StatusBadRequest)
	}
	userRegistration, err := authUtils.SerializeCredentials(&formBuffer)
	if err != nil && err.Error() != authUtils.NewValidationError(authUtils.EmptyPassword).Error() {
		return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	userInDb, err := models.GetUser(userRegistration)
	if err != nil && err != mongo.ErrNoDocuments {
		return utils.NewResponseError("Email not registered", http.StatusNotFound)
	}

	resetToken, err := token.IssueToken(&userInDb.ID, token.PasswordReset)
	if err != nil {
		return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	mailOrder := mailer.NewMailerOrder(
		userRegistration.Email,
		mailer.RecoverPasswordMail,
		&map[string]string{
			"resetToken": *resetToken,
		},
	)
	m.IssueMail(*mailOrder)

	fmt.Fprint(w, "Password reset request issued, please verify your email to proceed!\n")
	return nil
}

func ResetPasswordHandle(w http.ResponseWriter, r *http.Request) *utils.ResponseError {
	if r.Method != http.MethodPost {
		return utils.NewResponseError("Invalid method", http.StatusMethodNotAllowed)
	}

	ct := r.Header.Get("Content-Type")
	if ct == "" {
		return utils.NewResponseError("Bad Content-Type", http.StatusBadRequest)
	}

	resetToken := r.URL.Query().Get("token")
	if resetToken == "" {
		return utils.NewResponseError("Token cannot be nil", http.StatusBadRequest)
	}

	userId, err := token.ValidateToken(resetToken)
	if err != nil {
		if _, ok := err.(token.InvaildTokenError); ok {
			err = token.NewTokenValidationError(token.InvalidToken)
		}
		return utils.NewResponseError(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	formBuffer, err := io.ReadAll(r.Body)
	if err != nil {
		return utils.NewResponseError(err.Error(), http.StatusBadRequest)
	}
	userRegistration, err := authUtils.SerializeCredentials(&formBuffer)
	if err != nil {
		return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	err = models.UpdateUserPassword(*userId, userRegistration)
	if err != nil {
		return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Senha alterada com sucesso")
	return nil
}
