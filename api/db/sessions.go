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

	stmt := bson.D{primitive.E{Key: "id", Value: sessionId}}
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

	stmt := bson.D{primitive.E{Key: "id", Value: sessionCookie.Id}}
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

	cookies, err := GetCollection("cookies")
	if err != nil {
		log.Printf("Error on db connection: %v", err.Error())
		return "", err
	}

    newCookie, err := utils.GetRandomCookie()
	if err != nil {
		log.Printf("Error on cookie generation: %v", err.Error())
		return "", err
	}

	stmt := bson.D{
        primitive.E{Key: "email", Value: email},
        primitive.E{Key: "created_at", Value: time.Now()},
		// primitive.E{Key: "expire_at", Value: (time.Now().Add(5 * time.Minute))},
		primitive.E{Key: "expire_at", Value: (time.Now().Add(10 * time.Second))},
		primitive.E{Key: "id", Value: newCookie},
	}

    _, err = cookies.InsertOne(ctx, stmt)
	if err != nil {
		log.Printf("Error on saving new cookie: %v", err.Error())
		return "", err
	}

    return newCookie, nil
}
