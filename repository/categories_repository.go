package repository

import "github.com/LordRadamanthys/landing-page-manager/domain/categories"

type categoriesRepositoryInterface interface {
	Create(categories.Categories) error
	Get(string) (*categories.Categories, error)
}

type categoriesRepository struct{}

var (
	CategoriesRepository categoriesRepositoryInterface = &categoriesRepository{}
)

func (c *categoriesRepository) Create(category categories.Categories) error {
	return nil
}

func (c *categoriesRepository) Get(id string) (*categories.Categories, error) {

	return nil, nil
}
