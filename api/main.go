package main

import (
	"chat/api/db"
	"chat/api/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	log.Println("Starting things up...")
	log.Println("Listening on port 8080")

	http.HandleFunc("/login", makeHttpHandle(LoginUserHandle))
	http.HandleFunc("/register", makeHttpHandle(RegisterUserHandle))
	http.HandleFunc("/verify", makeHttpHandle(VerifySessionHandle))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Closing server: %v", err.Error())
	}
}

type HandleWithError func(http.ResponseWriter, *http.Request) error

func makeHttpHandle(f HandleWithError) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("Erro na resposta ao cliente: %v", err.Error())
		}
	}
}

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

	err = db.CreateUser(credentials)

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
		http.Error(w, "Erro na entrada das credenciais", http.StatusBadRequest)
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

	cookieId, err := db.CreateCookie(savedCredentials.Email)
	cookie := http.Cookie{
        Name: "sessionId",
        Value: cookieId,
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

func VerifySessionHandle(w http.ResponseWriter, r *http.Request) error {
	sessionId, err := r.Cookie("sessionId")
	if err != nil {
		http.Error(w, "Sessão não autenticada", http.StatusForbidden)
		return err
	}

	cookie, err := db.FindCookie(sessionId.Value)
	if err != nil {
		http.Error(w, "Sessão não autenticada", http.StatusForbidden)
		return err
	}

    if cookie.ExpireAt.Before(time.Now().UTC()) {
        err = db.DeleteCookie(cookie)
		http.Error(w, "Sessão não autenticada", http.StatusForbidden)
		return err
    }

    w.WriteHeader(http.StatusAccepted)
	return nil
}
