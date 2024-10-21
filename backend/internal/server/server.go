package server

import (
	"database/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Mux    http.Handler
	Server *http.Server
	Store  *sql.DB
}

func loggerToFile(filePath string) *log.Logger {
	fh, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("Error on logger creation: %s", err.Error())
	}

	return log.New(fh, "Chat-Server", log.LstdFlags|log.Lshortfile)
}

func NewApp(port, stdout string) *App {
	l := loggerToFile(stdout)

	mu := http.NewServeMux()
	serv := NewServer(port, mu, l)

	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		l.Fatalf("Error opening database: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		l.Fatalf("Error pinging database: %v", err.Error())
	}

	l.Println("Success on ping, db connection working!")

	return &App{
		Mux:    mu,
		Server: serv,
		Store:  db,
	}
}

func NewServer(port string, mux http.Handler, logger *log.Logger) *http.Server {
	server := http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 30,
		ErrorLog:     logger,
	}

	return &server
}
