package models

import (
	"chat/api/db"
	"chat/api/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRegistratinError struct {}
func (u UserRegistratinError)Error() string{
    return "Email already registered"
}

func CreateUser(cred *utils.Account) error{
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    users, err := db.GetCollection("users")
    if err != nil {
        return err
    } 

    stmt := bson.D{bson.E{Key: "email", Value: cred.Email}}
    count, err := users.CountDocuments(ctx, stmt)
    switch true {
    case err != nil:
        return err
    case count > 0:
        return UserRegistratinError{}
    }

    _, err = users.InsertOne(ctx, cred)
    return nil
}

