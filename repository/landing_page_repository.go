package repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	LandingPageRepository landingPageRepositoryInterface = &landingPageRepository{}
)

type landingPageRepository struct{}

type landingPageRepositoryInterface interface {
	InsertLandingPage(landingPage landing_page.LandingPage) error
	Update(landingPage landing_page.LandingPage) error
	GetTemplate(string, int) (*landing_page.LandingPage, error)
}

func (lp *landingPageRepository) InsertLandingPage(landingPage landing_page.LandingPage) error {

	landingPage.BrokersList = make([]brokers.Brokers, 0)
	result, err := db.GetLandingPageCollection().InsertOne(context.TODO(), landingPage)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (lp *landingPageRepository) Update(landingPage landing_page.LandingPage) error {
	idHex, _ := primitive.ObjectIDFromHex(landingPage.Id)
	filter := bson.M{"_id": idHex}
	update := bson.M{
		"$set": bson.M{
			"title":             "Carnaval",
			"template":          "1",
			"field_title":       "teste",
			"field_description": "test",
			"field_image":       "image"}}
	result, err := db.GetLandingPageCollection().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (lp *landingPageRepository) GetTemplate(idLP string, idBroker int) (*landing_page.LandingPage, error) {

	objectIdLP, _ := primitive.ObjectIDFromHex(idLP)
	fmt.Println(idLP, idBroker)
	filter := bson.M{"brokerslist": bson.M{"$elemMatch": bson.M{"id_broker": idBroker}, "_id": objectIdLP}}
	// filter := bson.M{"_id": objectIdLP, "brokerslist.id_broker": idBroker}
	obj := db.GetLandingPageCollection().FindOne(context.TODO(), filter)

	landingPage := &landing_page.LandingPage{}
	obj.Decode(landingPage)
	fmt.Println(landingPage)
	return landingPage, nil
	// fmt.Println(result)
}
