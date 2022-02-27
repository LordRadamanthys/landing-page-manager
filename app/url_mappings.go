package app

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/controllers"
	"github.com/gin-gonic/gin"
)

func loadUrls() {
	router.GET("/ping", pong)

	router.GET("broker/:id_broker/templates", controllers.BrokerController.GetTemplatesBroker)
	router.PATCH("broker/:id_template", controllers.BrokerController.UpdateBroker)
	router.POST("broker/", controllers.BrokerController.InsertBroker)
	router.POST("broker/upload", controllers.LandingPageController.UploadImage)

	router.GET("landingPage/:hash", controllers.LandingPageController.GetTemplate)
	router.POST("landingPage", controllers.LandingPageController.InsertLandingPage)

	router.POST("categories", controllers.CategoriesController.Create)
	router.GET("categories/:id", controllers.CategoriesController.Get)
	router.GET("categories", controllers.CategoriesController.GetAll)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong!")
}
