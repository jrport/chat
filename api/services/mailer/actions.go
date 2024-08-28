package mailer

import (
	"chat/api/utils"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

func (e *EmailService) VerifyAccount(user *Job) {
	rawToken := make([]byte, 32)
	_, err := rand.Reader.Read(rawToken)
	if err != nil {
		slog.Error("Error on token generation")
	}
	token := hex.EncodeToString(rawToken)
    slog.Info("id")

	rdClient := redis.NewClient(&redis.Options{Addr: ":6379"})
	err = rdClient.Set(
		context.TODO(),
		token,
		user.Account.ID.Hex(),
		time.Minute*15,
	).Err()
	if err != nil {
		slog.Error(err.Error() + "on storing key")
	}

	verificationLink := "http://localhost:8080/verify?token=" + token
	body := fmt.Sprintf(
		"<div>Verifique sua conta do whatsapp2 %s! Clique <a href='%s'>aqui</a> para verificar</div>",
		user.Account.Email,
		verificationLink,
	)

	emailData := utils.NewEmail("Confirme sua conta do whatsapp2!", body)
	msg := utils.WriteEmail(emailData)
	e.SendEmail(user.Account.Email, msg)
    user.Result <- nil
}
