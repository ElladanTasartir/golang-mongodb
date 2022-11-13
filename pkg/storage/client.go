package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbClient struct {
	client *mongo.Client
}

func Connect(hostname string) *DbClient {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(hostname))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB")
	}

	log.Printf("Succesfully conncected to MongoDB @%s", hostname)

	return &DbClient{
		client,
	}
}

func (client *DbClient) CloseConnection() {
	client.CloseConnection()
}
