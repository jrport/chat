package main

import (
	"chat/api/handles"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Ltime)
	log.SetFlags(log.Ldate)

	handles.SetupRoutes()
}

func main(){
	log.Print("Booting server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Shutting down server...", err.Error())
	}
}
