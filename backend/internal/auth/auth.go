package auth

import (
	"net/http"
	"fmt"
)

func AuthenticateSession(r *http.Request) error {
	cks, err := r.Cookie("_session_id")
	if err != nil {
		return fmt.Errorf("Invalid session")
	}
}
