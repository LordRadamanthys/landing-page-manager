package services

import (
	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"github.com/LordRadamanthys/landing-page-manager/repository"
)

var (
	LandingPageService landinPageInterface = &landinPageService{}
)

type landinPageService struct{}
type landinPageInterface interface {
	Insert(landing_page.LandingPage) error
	Update(landing_page.LandingPage) error
	Get(string, int) (*landing_page.LandingPage, error)
}

func (lp *landinPageService) Insert(landingPage landing_page.LandingPage) error {
	return repository.LandingPageRepository.InsertLandingPage(landingPage)
}

func (lp *landinPageService) Update(landingPage landing_page.LandingPage) error {
	return repository.LandingPageRepository.Update(landingPage)
}

func (lp *landinPageService) Get(idLP string, idBroker int) (*landing_page.LandingPage, error) {
	return repository.LandingPageRepository.GetTemplate(idLP, idBroker)
}
