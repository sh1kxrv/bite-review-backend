package mongodb

import (
	"bitereview/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
}

func NewMongoInstance(c *config.Config) (*MongoInstance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(c.Database.ConnectionURL)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoInstance{
		Client: client,
	}, err
}

func (m *MongoInstance) GetCollection(collection string) *mongo.Collection {
	return m.Client.Database(config.C.Database.Name).Collection(collection)
}
