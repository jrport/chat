package db

import (
	"chat/api/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateUser(credentials *utils.Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

    credentials.Hash()

	collection, err := GetCollection("users")
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(ctx, credentials)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(credentials *utils.Credentials) (*utils.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
    defer cancel()

	collection, err := GetCollection("users")
	if err != nil {
		return nil, err
	}

	stmt := bson.D{primitive.E{Key: "email", Value: credentials.Email}}
    queryResult := collection.FindOne(ctx, stmt)
    _, err = queryResult.Raw()
    if err != nil {
        return nil, err
    }

    newCredentials := new(utils.Credentials)
    err = queryResult.Decode(newCredentials)
    if err != nil {
        return nil, err
    }

    return newCredentials, nil
}


