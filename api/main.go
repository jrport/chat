package main

import (
	"chat/api/auth"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", MakeHandle(LoginRoute))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicf("Error forced closure of server: %v", err.Error())
	}
}

type HandleWithError func(http.ResponseWriter, *http.Request) error
type RouteHandle func(http.ResponseWriter, *http.Request)

func MakeHandle(f HandleWithError) RouteHandle{
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("Error when processing request: %s", err.Error())
		}
	}
}

func LoginRoute(w http.ResponseWriter, r *http.Request) error{
	w.Header().Set("Access-Control-Allow-Origin", "*") 
    cookie, error := auth.AutenticateLogin(&r.Body)
    http.SetCookie(w, cookie)
    if error != nil {
        log.Fatalf("Error: %v", error.Error())
    }
	w.Write([]byte("oi"))
	return nil
}

