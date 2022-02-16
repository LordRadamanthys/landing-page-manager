package services

import (
	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
	"github.com/LordRadamanthys/landing-page-manager/domain/template"
	"github.com/LordRadamanthys/landing-page-manager/repository"
)

var (
	BrokerService brokerInterface = &brokerService{}
)

type brokerService struct{}
type brokerInterface interface {
	Insert(string, brokers.Brokers)
	Update(string, brokers.Brokers)
	Get(string, string)
	GetTemplates(id string) *[]template.Template
}

func (b *brokerService) Insert(idLandiPage string, broker brokers.Brokers) {
	repository.BrokerRepository.InsertBroker(idLandiPage, broker)
}

func (b *brokerService) Update(idLandiPage string, broker brokers.Brokers) {
	repository.BrokerRepository.Update(idLandiPage, broker)
}

func (b *brokerService) Get(idLandiPage string, idBroker string) {

}

func (b *brokerService) GetTemplates(id string) *[]template.Template {
	return repository.BrokerRepository.GetAllTemplates(id)
}
