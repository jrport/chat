package main

import (
	"chat/backend/internal/server"
	"fmt"
)

func main() {
	server := server.NewServer(":8080")
	println("Server running on 8080...")
	if err := server.Run(); err != nil {
		fmt.Printf("Error on server: %v", err.Error())
		return
	}
	println("Server closing...")
	return
}
