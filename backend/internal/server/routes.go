package server

import (
	"fmt"
	"net/http"
)

func SetupRoutes() (m *http.ServeMux){
	m = http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "blabal")
	})

	return
}
