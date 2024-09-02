package models

import (
	"chat/api/database"
	"chat/api/handles/utils"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(userAccount *utils.UserRegistration) (*primitive.ObjectID, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
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

func GetUser(userAccount *utils.UserRegistration) (*utils.UserRegistration, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users, err := database.GetCollection("users")
	if err != nil {
		return nil, err
	}
	
	var userInDb utils.UserRegistration
	log.Printf("b: %v", userAccount)
	filter := bson.D{bson.E{Key: "email", Value: (*userAccount).Email}}
	err = users.FindOne(ctx, filter).Decode(userInDb)
	
	if err != nil {
		return nil, err
	}

	return nil, nil
}

