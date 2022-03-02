package services

import (
	"errors"
	"time"

	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/domain/landing_page"
	"github.com/LordRadamanthys/landing-page-manager/domain/template"
	"github.com/LordRadamanthys/landing-page-manager/repository"
	"github.com/LordRadamanthys/landing-page-manager/utils/base64_util"
)

var (
	LandingPageService landinPageInterface = &landinPageService{}
)

type landinPageService struct{}
type landinPageInterface interface {
	Insert(landing_page.LandingPage) error
	Update(landing_page.LandingPage) error
	Get(string) (*landing_page.LandingPage, error)
	ListAllLandingPages() (*[]template.Template, error)
}

func (lp *landinPageService) Insert(landingPage landing_page.LandingPage) error {
	landingPage.Date_Created = time.Now().Format("02/01/2006")
	landingPage.Brokers_list = make([]brokers.Brokers, 0)
	return repository.LandingPageRepository.InsertLandingPage(landingPage)
}

func (lp *landinPageService) Update(landingPage landing_page.LandingPage) error {
	return repository.LandingPageRepository.Update(landingPage)
}

func (lp *landinPageService) Get(hash string) (*landing_page.LandingPage, error) {
	idLP, idBroker, err := base64_util.DecodeUrlLandingPage(hash)
	if err != nil {
		return nil, err
	}
	return repository.LandingPageRepository.GetTemplate(idLP, idBroker)
}

func (lp *landinPageService) ListAllLandingPages() (*[]template.Template, error) {
	list, err := repository.LandingPageRepository.ListAllLandingPages()
	if err != nil {
		return nil, err
	}
	if len(*list) == 0 {
		return nil, errors.New("empty list")
	}
	return list, nil
}
