package main

import (
	"fmt"
	"jport/chat/backend/internal/db"
	"jport/chat/backend/internal/handlers"
	logging "jport/chat/backend/internal/logger"
	"jport/chat/backend/internal/server"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		panic("Please provide a port, the default logging destination (a json file)" +
			" and the sqlite database for the application.")
	}

	logger, err := logging.SetupLogger(os.Args[2])
	if err != nil {
		print("Error on logger setup.\n")
		panic(err.Error())
	}

	dbConn, err := db.SetupDb(os.Args[3])
	if err != nil {
		print("Error on database setup.\n")
		panic(err.Error())
	}

	app := server.NewApp(os.Args[1], logger, dbConn)
	
	handlers.SetupRoutes(app)
	if err := app.Run(); err != nil {
		app.Logger.Error(fmt.Sprintf("Closing with error: %s", err.Error()))
	}
}
