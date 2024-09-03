package token

import (
	"chat/api/redis"
	"context"
	"math/rand"
	"strings"
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

const (
	tokenLength = 24
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateToken() (*string, error){
    var sb strings.Builder
    sb.Grow(tokenLength)

    for i := 0; i < tokenLength; i++ {
        // Pick a random character from the characters string
        randomIndex := rand.Intn(len(characters))
        sb.WriteByte(characters[randomIndex])
    }
	strToken := sb.String()

	return &strToken, nil
}

func ValidateToken(token string) (*string, error){
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    client, err := redis.GetClient()
    if err != nil {
        return nil, err
    }
    
    userId, err := client.Get(ctx, token).Result()
    if err != nil {
        return nil, err
    }

    return &userId, nil
}

func ExpireToken(token string) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    client, err := redis.GetClient()
    if err != nil {
        return err
    }
    
    err = client.Expire(ctx, token, time.Second * 0).Err()
    if err != nil {
        return err
    }

    return nil
}
