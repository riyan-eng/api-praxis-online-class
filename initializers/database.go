package initializers

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectToDatabase() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	db = client.Database("praxis-online-class")
	fmt.Println("Success connect to database")
	return nil
}

func Collection(col string) *mongo.Collection {
	return db.Collection(col)
}

func CloseDatabase() error {
	return db.Client().Disconnect(context.Background())
}
