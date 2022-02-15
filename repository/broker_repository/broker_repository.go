package broker_repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	landingPageCollection = db.GetConnection().Database("db_landing_page").Collection("lp_collection")
)

func InsertBroker(idLP string, broker brokers.Brokers) {
	id, err := primitive.ObjectIDFromHex(idLP)
	if err != nil {
		panic(err)
	}
	update := bson.M{"$push": bson.M{"brokerslist": broker}}
	result, err := landingPageCollection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		panic(err)
	}
	if result.MatchedCount == 0 {
		panic("id n encontrado")
	}
	fmt.Println(result)
	fmt.Println(result.MatchedCount)
}

func Update(idLP string, broker brokers.Brokers) {
	idHex, _ := primitive.ObjectIDFromHex(idLP)
	filter := bson.M{"_id": idHex, "brokerslist.id": broker.Id}
	update := bson.M{"$set": bson.M{"brokerslist.$": broker}}
	result, err := landingPageCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func GetAllTemplates(id string) {
	filter := bson.M{"brokerslist.id": id}
	// update := bson.M{"$set": bson.M{"brokerslist.$": broker}}
	cursor, err := landingPageCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}
	// fmt.Println(result)
}
