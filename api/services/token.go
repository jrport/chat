package services

import (
	"chat/api/redis"
	"context"
	"crypto/rand"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IssueToken(uid *primitive.ObjectID) error{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	rdbClient, err := redis.GetClient()
	if err != nil {
		return err
	}

	token, err := generateToken()
	if err != nil {
		return err
	}
	
	uidHexed := uid.Hex()
	err = rdbClient.Set(ctx, *token, uidHexed, time.Minute * 10).Err()
	if err != nil {
		return err
	}
		
	return nil
}

func generateToken() (*string, error){
	randToken := make([]byte, 16)
	_, err := rand.Read(randToken)
	if err != nil {
		return nil, err
	}
	strToken := string(randToken[:])
	slog.Info("strtoken is " + strToken)
	return &strToken, nil
}
