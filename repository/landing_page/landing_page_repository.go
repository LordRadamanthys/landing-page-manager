package landing_page

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	landingPageCollection = db.GetConnection().Database("db_landing_page").Collection("lp_collection")
)

func InsertLandingPage(lp landing_page.LandingPage) {

	result, err := landingPageCollection.InsertOne(context.TODO(), lp)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Update(lp landing_page.LandingPage) {
	idHex, _ := primitive.ObjectIDFromHex(lp.Id)
	filter := bson.M{"_id": idHex}
	update := bson.M{
		"$set": bson.M{
			"title":             "Carnaval",
			"template":          "1",
			"field_title":       "teste",
			"field_description": "test",
			"field_image":       "image"}}
	result, err := landingPageCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
