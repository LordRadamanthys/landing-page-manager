package controllers

import (
	"net/http"

	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"github.com/LordRadamanthys/landing-page-manager/services"
	"github.com/LordRadamanthys/landing-page-manager/utils/base64_util"
	"github.com/gin-gonic/gin"
)

type landingPageControllerInterface interface {
	InsertLandingPage(*gin.Context)
	Update(*gin.Context)
	UploadImage(c *gin.Context)
	GetTemplate(c *gin.Context)
	ListAllLandingPages(c *gin.Context)
}

type landingPageController struct{}

const (
	access_control_allow_origin = "Access-Control-Allow-Origin"
)

var (
	LandingPageController landingPageControllerInterface = &landingPageController{}
)

func (lp *landingPageController) InsertLandingPage(c *gin.Context) {
	c.Writer.Header().Set(access_control_allow_origin, "*")
	var landing landing_page.LandingPage

	if err := c.ShouldBindJSON(&landing); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	if err := services.LandingPageService.Insert(landing); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (lp *landingPageController) Update(c *gin.Context) {
	c.Writer.Header().Set(access_control_allow_origin, "*")
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

func (lp *landingPageController) GetTemplate(c *gin.Context) {
	c.Writer.Header().Set(access_control_allow_origin, "*")
	hash := c.Param("hash")

	response, err := services.LandingPageService.Get(hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (lp *landingPageController) ListAllLandingPages(c *gin.Context) {
	c.Writer.Header().Set(access_control_allow_origin, "*")
	response, err := services.LandingPageService.ListAllLandingPages()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (lp *landingPageController) UploadImage(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	imgBase64Str, base64Err := base64_util.ConvertImageToBase64(file)
	if base64Err != nil {
		c.JSON(http.StatusBadRequest, base64Err)
		return
	}
	img2html := "<html><body><img width=\"200\" hegth=\"200\" src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"

	c.Writer.WriteString(img2html)
}
