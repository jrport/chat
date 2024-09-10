package mailer

import (
	"fmt"
	"log/slog"
	"net/smtp"
)

type MailAction int

const (
	ValidationTokenMail MailAction = iota
	RecoverPasswordMail
)

func (s *Mailer) SendMail(email, subject, body string) error {
	auth := s.GetAuth()
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	msg := []byte(subject + mime + body)

	err := smtp.SendMail(s.Config.Host+":"+s.Config.Port, auth, s.Config.Email, []string{email}, msg)
	if err != nil {
		slog.Error("Error " + err.Error() + " | Sending email to " + email)
		return err
	}

	slog.Info("Mail sent successfully to " + email)
	return nil
}

func (s *Mailer) SendValidationMail(email, token string) {
	subject := "Subject: Confirmation Token for WhatsappForum\r\n"

	url := fmt.Sprintf("http://localhost:8080/verify?token=%s", token)
	body := fmt.Sprintf(`
	<html>
	  <body>
	    <h1>Oii!</h1>
		Clique <a href='%s'>aqui</a> para confirmar seu cadastro no WhatsappForum!
		Se vc não sabe o que é isso só ignora, bjs XOXO
	  </body>
	</html>
	`, url,
	)

	if err := s.SendMail(email, subject, body); err != nil {
		slog.Error("Error sending verification email to " + email)
	}

	slog.Info("Email sent successfully to " + email)
}

func (s *Mailer) SendForgotPasswordMail(email string) {
}
