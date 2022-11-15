package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbClient struct {
	client   *mongo.Client
	database *mongo.Database
}

func Connect(hostname string, database string) *DbClient {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(hostname))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB")
	}

	log.Printf("Succesfully conncected to MongoDB @%s", hostname)

	return &DbClient{
		client:   mongoClient,
		database: mongoClient.Database(database),
	}
}

func (db *DbClient) CloseConnection() {
	err := db.client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
