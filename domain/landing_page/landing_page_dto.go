package landing_page

import (
	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LandingPage struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Id_mkt            int                `json:"id_mkt"`
	Id_category       string             `json:"id_category"`
	Title             string             `json:"title"`
	Template          string             `json:"template"`
	Field_title       string             `json:"field_title"`
	Field_description string             `json:"field_description"`
	Field_image       string             `json:"field_image"`
	Active            bool               `json:"active"`
	Date_Created      string             `json:"date_created"`
	Brokers_list      []brokers.Brokers  `json:"brokers_list"`
}
