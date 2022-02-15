package db

import (
	"context"

	"github.com/LordRadamanthys/landing-page-manager/utils/environment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(environment.GetEnvVariables("CONNECTION_MONGO")))
	if err != nil {
		panic(err)
	}

	return client
}
