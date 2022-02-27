package repository

import (
	"context"
	"fmt"

	"github.com/LordRadamanthys/landing-page-manager/db"
	"github.com/LordRadamanthys/landing-page-manager/domain/categories"
	"go.mongodb.org/mongo-driver/bson"
)

type categoriesRepositoryInterface interface {
	Create(categories.Categories) error
	Get(string) (*categories.Categories, error)
	GetAll() (*[]categories.Categories, error)
}

type categoriesRepository struct{}

var (
	CategoriesRepository categoriesRepositoryInterface = &categoriesRepository{}
)

func (c *categoriesRepository) Create(category categories.Categories) error {
	_, err := db.GetCategoriesCollection().InsertOne(context.TODO(), category)
	if err != nil {
		return err
	}
	return nil
}

func (c *categoriesRepository) Get(id string) (*categories.Categories, error) {
	result := db.GetCategoriesCollection().FindOne(context.TODO(), id)
	var category categories.Categories
	if err := result.Decode(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categoriesRepository) GetAll() (*[]categories.Categories, error) {

	cursor, err := db.GetCategoriesCollection().Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var category []categories.Categories

	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &category); err != nil {
		panic(err)
	}

	return &category, nil
}
