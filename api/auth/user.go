package auth

import (
	"chat/api/db"
	"fmt"
	"net/http"
	"strings"

	// "fmt"
	"io"
	"log"
	// "encoding/json"
	// "fmt"
	// "io"
	// "log"
	// "go.mongodb.org/mongo-driver/bson"
)

type AuthError struct{}

func (error AuthError) Error() string {
	return "Usuário inválido"
}

func AutenticateLogin(login *io.ReadCloser) (*http.Cookie, error) {
	cookie, err := db.AuthenticateUser(login)
	switch {
	case err != nil && strings.HasSuffix(err.Error(), "no documents in result"):
		return &http.Cookie{}, fmt.Errorf("Usuário Inválido")
	case err != nil:
		return nil, fmt.Errorf("Error na autenticação %s", err.Error())
	default:
		log.Print("User entrou!")
		return cookie, nil
		// cookie := db.GetCookie()
	}
}
