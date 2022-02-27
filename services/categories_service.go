package services

import (
	"github.com/LordRadamanthys/landing-page-manager/domain/categories"
	"github.com/LordRadamanthys/landing-page-manager/repository"
)

type categoriesInterface interface {
	GetAll() (*[]categories.Categories, error)
	Get(string) (*categories.Categories, error)
	Create(categories.Categories) error
}
type categoriesService struct{}

var CategoriesService categoriesInterface = &categoriesService{}

func (cat *categoriesService) GetAll() (*[]categories.Categories, error) {
	return repository.CategoriesRepository.GetAll()
}

func (cat *categoriesService) Get(id string) (*categories.Categories, error) {
	return repository.CategoriesRepository.Get(id)
}

func (cat *categoriesService) Create(category categories.Categories) error {
	return repository.CategoriesRepository.Create(category)
}
