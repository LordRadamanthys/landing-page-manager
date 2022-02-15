package services

import (
	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	repository "github.com/LordRadamanthys/landing-page-manager/repository/landing_page"
)

var (
	LandingPageService landinPageInterface = &landinPageService{}
)

type landinPageService struct{}
type landinPageInterface interface {
	Insert(landing_page.LandingPage)
	Update(landing_page.LandingPage)
	Get(string)
}

func (lp *landinPageService) Insert(landingPage landing_page.LandingPage) {
	repository.InsertLandingPage(landingPage)
}

func (lp *landinPageService) Update(landingPage landing_page.LandingPage) {
	repository.Update(landingPage)
}

func (lp *landinPageService) Get(id string) {

}
