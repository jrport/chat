package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI string = "mongodb://localhost:27017"

func GetClient() (*mongo.Client, error){ 
    timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(timeoutCtx, clientOpts)
    if err != nil {
        return nil, err
    }
    return client, nil
}

func GetCollection(collection string) (*mongo.Collection, error) {
    client, err := GetClient()
    if err != nil {
        return nil, err
    } 
    
    database := client.Database("chat")
    c := database.Collection(collection)

    return c, nil
}
