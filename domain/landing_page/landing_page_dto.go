package landing_page

import "github.com/LordRadamanthys/landing-page-manager/domain/brokers"

type LandingPage struct {
	Id                string            `json:"id"`
	Title             string            `json:"title"`
	Template          string            `json:"template"`
	Field_title       string            `json:"fieldTitle"`
	Field_description string            `json:"fieldDescription"`
	Field_image       string            `json:"fieldImage"`
	BrokersList       []brokers.Brokers `json:"brokers"`
}
