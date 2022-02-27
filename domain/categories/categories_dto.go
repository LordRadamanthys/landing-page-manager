package categories

import "go.mongodb.org/mongo-driver/bson/primitive"

type Categories struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title"`
}
