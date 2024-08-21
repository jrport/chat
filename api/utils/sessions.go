package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type Cookie struct {
	Id        string    `bson:"id"`
	User      string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
	ExpireAt  time.Time `bson:"expire_at"`
}

func GetRandomCookie() (string, error){
	b := make([]byte, 10)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString([]byte(b)), nil
}
