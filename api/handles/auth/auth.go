package handles

import (
	"chat/api/db"
	"chat/api/mailer"
	"chat/api/utils"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var EmailChan = make(chan string)

func RegisterUserHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Apenas POST'S são permitidos", http.StatusMethodNotAllowed)
	}

	credentials, err := utils.ReadJson(&r.Body)
	if err != nil {
		errorMsg := fmt.Sprintf("Não foi possível concluir o cadastro: %v", err.Error())
		http.Error(w, errorMsg, http.StatusNotFound)
		return err
	}

	log.Println("scheduling confirmation mail")
	err = db.CreateUser(credentials)
	EmailChan <- credentials.Email

	if err != nil {
		errorMsg := fmt.Sprintf("Não foi possível concluir o cadastro: %v", err.Error())
		http.Error(w, errorMsg, http.StatusNotFound)
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	log.Print("Registrado com sucesso")
	fmt.Fprint(w, "Registrado com sucesso")

	return nil
}

func LoginUserHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Apenas POST'S são permitidos", http.StatusMethodNotAllowed)
	}

	credentials, err := utils.ReadJson(&r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	savedCredentials, err := db.GetUser(credentials)
	if err != nil {
		http.Error(w, "Não foi encontrado o usuário", http.StatusForbidden)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(savedCredentials.PasswordHash), []byte(credentials.PasswordHash)); err != nil {
		http.Error(w, "Senha incorreta", http.StatusForbidden)
		return err
	}

	if !savedCredentials.Confirmed {
		http.Error(w, "Email não confirmado", http.StatusUnauthorized)
		return err
	}

	cookieId, err := db.CreateCookie(savedCredentials.Email)
	cookie := http.Cookie{
		Name:   "sessionId",
		Value:  cookieId,
		Quoted: false,
	}
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/plain")
	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
	fmt.Fprint(w, "Logado com sucesso")

	log.Println("Logado com sucesso")

	return nil
}

func LogOutHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		http.Error(w, "Apenas DELETE'S são permitidos", http.StatusMethodNotAllowed)
	}
	sessionId, err := r.Cookie("sessionId")
	if err != nil {
		http.Error(w, "Não está logado", http.StatusForbidden)
		return err
	}
	cookie, err := db.FindCookie(sessionId.Value)
	if err != nil {
		http.Error(w, "Cookie inválido", http.StatusForbidden)
		return err
	}
	if err = db.DeleteCookie(cookie); err != nil {
		http.Error(w, "Sessão não autenticada", http.StatusForbidden)
		return err
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "Sessão encerrada com sucesso")
	return nil
}

func PingRoute(w http.ResponseWriter, _ *http.Request) error {
	w.Write([]byte("bla"))
	return nil
}

func init() {
	mailerConfig := mailer.NewEmailConfiguration(
		utils.E.Email,
		utils.E.Password,
		utils.E.Host,
		utils.E.Port,
	)
	EmailService := mailer.NewEmailService(mailerConfig)

	go EmailService.Run(EmailChan)
}
