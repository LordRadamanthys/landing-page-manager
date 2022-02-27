package controllers

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/domain/categories"
	"github.com/LordRadamanthys/landing-page-manager/services"
	"github.com/gin-gonic/gin"
)

type categoriesControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
}

type categoriesController struct{}

var (
	CategoriesController categoriesControllerInterface = &categoriesController{}
)

func (cat *categoriesController) Create(c *gin.Context) {
	var category categories.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	createErr := services.CategoriesService.Create(category)
	if createErr != nil {
		c.JSON(http.StatusBadRequest, createErr)
		return
	}
	c.Status(http.StatusCreated)
}

func (cat *categoriesController) Get(c *gin.Context) {
	id := c.Param("id")

	category, err := services.CategoriesService.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, category)
}

func (cat *categoriesController) GetAll(c *gin.Context) {

	categories, err := services.CategoriesService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, categories)
}
