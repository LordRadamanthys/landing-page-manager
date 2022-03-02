package controllers

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/services"
	"github.com/gin-gonic/gin"
)

type brokerControllerInterface interface {
	InsertBroker(c *gin.Context)
	UpdateBroker(c *gin.Context)
	GetTemplatesBroker(c *gin.Context)
}

type brokerController struct{}

var (
	BrokerController brokerControllerInterface = &brokerController{}
)

func (b *brokerController) InsertBroker(c *gin.Context) {
	var broker brokers.Brokers

	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
	}
	services.BrokerService.Insert("621f6edc19950530c43db8aa", broker)
}

func (b *brokerController) GetTemplatesBroker(c *gin.Context) {
	id_broker := c.Param("id_broker")
	response, err := services.BrokerService.GetTemplates(id_broker)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (b *brokerController) UpdateBroker(c *gin.Context) {
	var broker brokers.Brokers
	id_lp := c.Param("id_template")
	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
	}
	services.BrokerService.Update(id_lp, broker)
	c.JSON(http.StatusOK, nil)
}
