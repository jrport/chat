package main

import (
	"fmt"
	"jport/chat/backend/internal/server"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		panic("Please provide a port, the default logging destination (a json file)" +
			" and the sqlite database for the application.")
	}

	app := server.NewApp(os.Args[1], os.Args[2], os.Args[3])
	if err := app.Run(); err != nil {
		app.Logger.Error(fmt.Sprintf("Closing with error: %s", err.Error()))
	}
}
