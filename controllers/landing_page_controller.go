package controllers

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"github.com/LordRadamanthys/landing-page-manager/services"
	"github.com/gin-gonic/gin"
)

type landingPageControllerInterface interface {
	InsertLandingPage(*gin.Context)
	Update(*gin.Context)
}

type landingPageController struct{}

var (
	LandingPageController landingPageControllerInterface = &landingPageController{}
)

func (lp *landingPageController) InsertLandingPage(c *gin.Context) {
	var landing landing_page.LandingPage

	if err := c.ShouldBindJSON(&landing); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	if err := services.LandingPageService.Insert(landing); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (lp *landingPageController) Update(c *gin.Context) {
	var landing landing_page.LandingPage

	if err := c.ShouldBindJSON(&landing); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	if err := services.LandingPageService.Update(landing); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
