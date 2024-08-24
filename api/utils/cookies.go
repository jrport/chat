package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cookie struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Value     string             `bson:"value"`
	UserId    primitive.ObjectID `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	ExpireAt  time.Time          `bson:"expire_at"`
}

func GetRandomCookie() (string, error) {
	b := make([]byte, 10)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString([]byte(b)), nil
}
