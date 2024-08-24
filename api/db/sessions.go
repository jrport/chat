package db

import (
	"chat/api/utils"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindCookie(sessionId string) (*utils.Cookie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cookies, err := GetCollection("cookies")
	if err != nil {
		log.Printf("Error on db connection: %v", err.Error())
		return nil, err
	}

	stmt := bson.D{primitive.E{Key: "value", Value: sessionId}}
	queryResult := cookies.FindOne(ctx, stmt)
	if err := queryResult.Err(); err != nil {
		return nil, err
	}
	cookie := new(utils.Cookie)
	err = queryResult.Decode(cookie)
	if err != nil {
		log.Printf("Error on cookie decoding: %v", err.Error())
		return nil, err
	}

	return cookie, nil
}

func DeleteCookie(sessionCookie *utils.Cookie) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cookies, err := GetCollection("cookies")
	if err != nil {
		log.Printf("Error on db connection: %v", err.Error())
		return err
	}

	stmt := bson.D{primitive.E{Key: "_id", Value: sessionCookie.ID}}
	_, err = cookies.DeleteOne(ctx, stmt)
	if err != nil {
		log.Printf("Error on cookie deletion: %v", err.Error())
		return err
	}

	return nil
}

func CreateCookie(email string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

    users, err := GetCollection("users")
	if err != nil {
		log.Printf("Error on db connection: %v", err.Error())
		return "", err
	}

    user := utils.Credentials{}
    userQuery := users.FindOne(ctx, bson.D{bson.E{Key: "email", Value: email}})
    err = userQuery.Decode(&user)
    if err != nil {
		log.Printf("User not found on session generation: %v", err.Error())
		return "", err
    }

	cookies, err := GetCollection("cookies")
	if err != nil {
		log.Printf("Error on db connection: %v", err.Error())
		return "", err
	}

    newSessionID, err := utils.GetRandomCookie()
	if err != nil {
		log.Printf("Error on cookie generation: %v", err.Error())
		return "", err
	}
    newCookie := utils.Cookie{
        UserId: user.ID,
        Value: newSessionID,
        ExpireAt: time.Now().Add(10 * time.Minute),
        CreatedAt: time.Now(),
    }

    _, err = cookies.InsertOne(ctx, newCookie)
	if err != nil {
		log.Printf("Error on saving new cookie: %v", err.Error())
		return "", err
	}

    return newSessionID, nil
}
