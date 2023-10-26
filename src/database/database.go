package database

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mailman/src/config"
)

var client *mongo.Client

func Connect() {
	var err error
	opts := options.Client().ApplyURI(config.Env.MongoConnectionString).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	log.Info("Connected to MongoDB successfully!")
}

func Disconnect() error {
	err := client.Disconnect(context.Background())
	return err
}

func Use(databaseName string) *mongo.Database {
	return client.Database(databaseName)
}

func UseDefault() *mongo.Database {
	return client.Database("mailman")
}
