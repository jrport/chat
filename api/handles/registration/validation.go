package registration

import (
	"chat/api/database/models"
	"chat/api/handles/utils"
    tokenService"chat/api/services/token"
	"fmt"
	"net/http"
)

func ValidationHandle(w http.ResponseWriter, r *http.Request) *utils.ResponseError {
    if r.Method != http.MethodGet {
        return utils.NewResponseError("Invalid Method", http.StatusMethodNotAllowed)
    }
    
    userId := r.URL.Query().Get("id")
    token := r.URL.Query().Get("token")
    
    if (userId == "" || token == "") {
        return utils.NewResponseError("Invalid parameters", http.StatusBadRequest)
    }
    
    storedId, err := tokenService.ValidateToken(token)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }

    if userId != *storedId {
        return utils.NewResponseError("Invalid token", http.StatusBadRequest)
    }

    err = models.ValidateUser(userId)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }

    err = tokenService.ExpireToken(token)
    if err != nil {
        return utils.NewResponseError(err.Error(), http.StatusInternalServerError)
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "User validated successfully")
    return nil
}
