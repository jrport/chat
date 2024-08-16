package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI string = "mongodb://jrporto:123@localhost/?retryWrites=true&w=majority"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type UserCredentials struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `json:"username" bson:"login"`
	Password string             `json:"password" bson:"password"`
}

func requestToCredentials(login *io.ReadCloser) (*UserCredentials, error) {
	var credentials UserCredentials
	err := json.NewDecoder(*login).Decode(&credentials)
	if err != nil {
		return nil, err
	}
	return &credentials, nil
}

func credentialsToFilter(credentials UserCredentials) *bson.D {
	return &bson.D{
		{Key: "login", Value: credentials.Username},
		{Key: "password", Value: credentials.Password},
	}
}

func newClient() (*mongo.Client, error) {
	ctx := context.TODO()
	ServerApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(ServerApi)
	client, error := mongo.Connect(ctx, opts)
	if error != nil {
		log.Printf("Erro na conexão com o banco: %v", error.Error())
		return nil, error
	}
	return client, nil
}

func GetCollection(collection string) *mongo.Collection {
	client, _ := newClient()
	return client.Database("users").Collection("accounts")
}

func AuthenticateUser(login *io.ReadCloser) (*http.Cookie, error) {
	credentials, error := requestToCredentials(login)
	if error != nil {
		log.Printf("Payload de autenticação inválido, erro: %v", error.Error())
	}
	ctx := context.TODO()
	collection := GetCollection("users")
	filter := credentialsToFilter(*credentials)
	queryResult := collection.FindOne(ctx, filter)

	err := queryResult.Decode(&credentials)
	// log.Printf("No bson %s e %s | %v", credentials.Username, credentials.Password, credentials.Id)
	if err != nil {
		log.Printf("Error na query: %s", err.Error())
		return nil, err
	}
	cookie, err := getNewCookie(credentials.Id)
    log.Printf("Biscoite pro %s pronto!", credentials.Username)
	return cookie, err
}

func getNewCookie(id primitive.ObjectID) (*http.Cookie, error) {
	ctx := context.TODO()
	cookieValue := make([]rune, 30)
	for i := range cookieValue {
		cookieValue[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	filter := bson.D{{Key: "_id", Value: id}}
	stmt := bson.D{{Key: "$set", Value: bson.D{{Key: "cookie", Value: cookieValue}}}}
	collection := GetCollection("users")
	// count, error := collection.UpdateByID(ctx, id, stmt)
	count, error := collection.UpdateOne(ctx, filter, stmt)
	switch {
	case error != nil:
		log.Printf("Error on cookie generation %s", error.Error())
		return nil, error
	case count.MatchedCount == 0:
		log.Printf("No cookies generated for user id: %v", id)
		return nil, fmt.Errorf("Error on cookie generation")
	}
	log.Printf("setting user cookie: %v", cookieValue)
	return &http.Cookie{Name: "sessionId", Value: string(cookieValue)}, nil
}
