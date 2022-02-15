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
	InsertLandingPage(landingPage landing_page.LandingPage)
	Update(landingPage landing_page.LandingPage)
}

func (lp *landingPageRepository) InsertLandingPage(landingPage landing_page.LandingPage) {

	result, err := db.GetLandingPageCollection().InsertOne(context.TODO(), landingPage)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func (lp *landingPageRepository) Update(landingPage landing_page.LandingPage) {
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
		panic(err)
	}
	fmt.Println(result)
}
