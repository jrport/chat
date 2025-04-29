package hash

import (
	"os"

	"golang.org/x/crypto/argon2"
)

var (
	Salt        []byte
	Iterations  uint32 = 3
	Parellilism uint8  = 1
	Memory      uint32 = 64 * 1024
	Length      uint32 = 32
)

func init() {
	envSalt := os.Getenv("HASH_SALT")
	if envSalt == "" {
		panic("Please set the env variable HASH_SALT")
	}
	Salt = []byte(envSalt)
}

func HashPassword(password string) string{
	hash := argon2.IDKey([]byte(password), []byte(Salt), Iterations, Memory, Parellilism, Length)
	return string(hash)
}
