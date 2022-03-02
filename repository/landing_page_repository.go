package repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
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

	result, err := db.GetLandingPageCollection().InsertOne(context.TODO(), landingPage)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func (lp *landingPageRepository) Update(landingPage landing_page.LandingPage) error {
	idHex, _ := primitive.ObjectIDFromHex(landingPage.Id.String())
	filter := bson.M{"_id": idHex}
	update := bson.M{
		"$set": bson.M{
			"title":             landingPage.Title,
			"template":          landingPage.Template,
			"field_title":       landingPage.Field_title,
			"field_description": landingPage.Field_description,
			"field_image":       landingPage.Field_image,
			"active":            landingPage.Active,
		}}
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
	filter := bson.M{"brokers_list": bson.M{"$elemMatch": bson.M{"id_broker": idBroker}}, "_id": objectIdLP}
	obj := db.GetLandingPageCollection().FindOne(context.TODO(), filter)
	landingPage := &landing_page.LandingPage{}
	obj.Decode(landingPage)
	for _, broker := range landingPage.Brokers_list {
		if broker.Id_broker == idBroker {
			landingPage.Brokers_list = nil
			landingPage.Brokers_list = append(landingPage.Brokers_list, broker)
			break
		}
	}
	return landingPage, nil
}
