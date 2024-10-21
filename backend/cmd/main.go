package main

import (
	"log"
	"os"

	"chat.backend/internal/server"
)

func main(){
	if len(os.Args) != 2 {
		log.Fatal("Please provide the desired port to bind the server")
	}

	_ = server.NewApp(os.Args[1], "history.log")
}
