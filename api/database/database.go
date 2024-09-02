package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collectionName string) (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	collection := client.Database("chat").Collection(collectionName)

	return collection, nil
}


type DatabaseError struct{
	Msg string
}

func (d DatabaseError)Error() string{
	return d.Msg
}
	
func NewDatabaseError(errorMsg string) *DatabaseError {
	return &DatabaseError{
		Msg: errorMsg,
	}
}
