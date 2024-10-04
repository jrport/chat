package models

import (
	"chat/api/database"
	"chat/api/handles/auth/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(userAccount *utils.UserRegistration) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
	if err != nil {
		return nil, err
	}

	err = userAccount.Hash()
	if err != nil {
		return nil, err
	}

	result, err := users.InsertOne(ctx, userAccount)
	if err != nil {
		return nil, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, database.NewDatabaseError(fmt.Sprintf("Unexpected error on insert conversion of %v", id))
	}

	return &id, nil
}

func GetUser(userAccount *utils.UserRegistration) (*utils.UserRegistration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
	if err != nil {
		return nil, err
	}

	var userInDb utils.UserRegistration
	filter := bson.D{bson.E{Key: "Email", Value: userAccount.Email}}
	err = users.FindOne(ctx, filter).Decode(&userInDb)

	if err != nil {
		return nil, err
	}

	return &userInDb, nil
}

func ValidateUser(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.M{"$set": bson.M{"Confirmed": true}}
	err = users.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserPassword(userId string, userAccount *utils.UserRegistration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
	if err != nil {
		return err
	}

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	err = userAccount.Hash()
	if err != nil {
		return err
	}

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.M{"$set": bson.M{"Password": userAccount.Password}}
	err = users.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		return err
	}

	return nil
}
