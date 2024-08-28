package models

import (
	"chat/api/db"
	"chat/api/utils"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRegistratinError struct{}

func (u UserRegistratinError) Error() string {
	return "Email already registered"
}

func CreateUser(cred *utils.Account) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	users, err := db.GetCollection("users")
	if err != nil {
		return nil, err
	}

	stmt := bson.D{bson.E{Key: "email", Value: cred.Email}}
	var userInDb utils.Account
	err = users.FindOne(ctx, stmt).Decode(&userInDb)
	switch true {
	case err == mongo.ErrNoDocuments:
		return users.InsertOne(ctx, cred)
	case err != nil:
		return nil, err
	}

	if userInDb.Confirmed {
		return nil, UserRegistratinError{}
	}
	return &mongo.InsertOneResult{InsertedID: userInDb.ID}, nil
}

type InvalidTokenErr struct{}

func (i InvalidTokenErr) Error() string {
	return "Expired or forged token"
}

type AlreadyValidatedError struct {}

func (a AlreadyValidatedError) Error() string {
    return "Already validated User"
}

func ValidateUser(userToken string) error {
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	rdClient := redis.NewClient(&redis.Options{Addr: ":6379"})
	userId, err := rdClient.Get(context, userToken).Result()
	if err == redis.Nil {
		return InvalidTokenErr{}
	}

	users, err := db.GetCollection("users")
	if err != nil {
		return err
	}

    var updatedAccount utils.Account
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

    filter := bson.D{primitive.E{Key: "_id", Value: id}, primitive.E{Key: "confirmed", Value: false}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "confirmed", Value: true}}}}
    err = users.FindOneAndUpdate(context, filter, update).Decode(&updatedAccount)
    if err == mongo.ErrNoDocuments {
        return AlreadyValidatedError{}
    }
    return err
}
