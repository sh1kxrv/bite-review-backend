package database

import (
	"bitereview/config"
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var MongoClient *mongo.Client

func InitMongo(c *config.Config) {
	/*
		.SetBSONOptions(&options.BSONOptions{
			NilMapAsEmpty:       true,
			NilSliceAsEmpty:     true,
			NilByteSliceAsEmpty: true,
		})
	*/
	clientOptions := options.Client().ApplyURI(c.Database.ConnectionURL)
	clientOptions.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatalf("Failed to connect to mongodb: %s", err.Error())
	}

	MongoClient = client
}

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database(config.C.Database.Name).Collection(collection)
}
