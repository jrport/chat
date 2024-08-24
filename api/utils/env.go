package utils

import (
	"log"

	"github.com/goloop/env"
)


type Settings struct {
    Email string `env:"Email"`
    Password string `env:"Password"`
}

var E Settings

func init() {
    log.Println("Loading environment variables...")
    if err := env.Update(".env"); err != nil {
        log.Println("Error on loading env file")
        log.Fatal(err)
    }

    if err := env.Unmarshal("", &E); err != nil {
        log.Println("Error on env file unmarshalling")
        log.Fatal(err)
    }
}
