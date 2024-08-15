package main

import (
	// "fmt"
	"encoding/json"
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
	log.Print("Responedndn")
	msg := userCredentials{}
	w.Header().Set("Access-Control-Allow-Origin", "*") 
	json.NewDecoder(r.Body).Decode(&msg)
	// s := len(msg.username)
	log.Printf("Msg from client: %s ", msg.Username)
	w.Write([]byte("oi"))
	return nil
}

type userCredentials struct {
	Username, Password string;
}
