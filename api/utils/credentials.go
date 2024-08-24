package utils

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"password" bson:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	Confirmed    bool               `bson:"confirmed"`
}

func (c *Credentials) Hash() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(c.PasswordHash), bcrypt.DefaultCost)
	c.PasswordHash = string(hashed)
}
