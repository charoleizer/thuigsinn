package internal

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() (*mongo.Database, error) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB: ", mongoURI)
	fmt.Println("Database: ", os.Getenv("MONGO_DATABASE"))
	return client.Database(os.Getenv("MONGO_DATABASE")), nil
}
