package repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/domain/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	BrokerRepository brokerRepositoryInterface = &brokerRepository{}
)

type brokerRepository struct{}

type brokerRepositoryInterface interface {
	InsertBroker(string, brokers.Brokers)
	Update(string, brokers.Brokers)
	GetAllTemplates(string) *[]template.Template
}

func (br *brokerRepository) InsertBroker(idLP string, broker brokers.Brokers) {
	id, err := primitive.ObjectIDFromHex(idLP)
	if err != nil {
		panic(err)
	}
	update := bson.M{"$push": bson.M{"brokerslist": broker}}
	result, err := db.GetLandingPageCollection().UpdateByID(context.TODO(), id, update)
	if err != nil {
		panic(err)
	}
	if result.MatchedCount == 0 {
		panic("id n encontrado")
	}
	fmt.Println(result)
	fmt.Println(result.MatchedCount)
}

func (br *brokerRepository) Update(idLP string, broker brokers.Brokers) {
	idHex, _ := primitive.ObjectIDFromHex(idLP)
	filter := bson.M{"_id": idHex, "brokerslist.id": broker.Id}
	update := bson.M{"$set": bson.M{"brokerslist.$": broker}}
	result, err := db.GetLandingPageCollection().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func (br *brokerRepository) GetAllTemplates(id string) *[]template.Template {
	filter := bson.M{"brokerslist.id": id}
	// update := bson.M{"$set": bson.M{"brokerslist.$": broker}}
	cursor, err := db.GetLandingPageCollection().Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var obj []template.Template
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &obj); err != nil {
		panic(err)
	}
	for _, result := range obj {

		fmt.Println(result)
	}
	return &obj
	// fmt.Println(result)
}
