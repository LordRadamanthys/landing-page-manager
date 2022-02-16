package app

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/controllers"
	"github.com/gin-gonic/gin"
)

func loadUrls() {
	router.GET("/ping", pong)

	router.GET("landingPage/broker/:id_broker/templates", controllers.BrokerController.GetTemplatesBroker)
	router.PATCH("landingPage/broker/:id_template", controllers.BrokerController.UpdateBroker)
	router.POST("landingPage/broker/", controllers.BrokerController.InsertBroker)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "Pong!")
}
