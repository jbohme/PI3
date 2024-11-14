package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URI      = "MONGODB_URI"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URI)
	mongodb_database := os.Getenv(MONGODB_DATABASE)

	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_url))
	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
