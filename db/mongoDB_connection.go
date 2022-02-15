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

func GetLandingPageCollection() *mongo.Collection {
	return GetConnection().Database("db_landing_page").Collection("lp_collection")
}
