package landing_page

import (
	"time"

	"github.com/LordRadamanthys/landing-page-manager/domain/brokers"
)

type LandingPage struct {
	Id                string            `json:"id"`
	Id_mkt            int               `json:"id_mkt"`
	Id_categoria      string            `json:"id_categoria"`
	Title             string            `json:"title"`
	Template          string            `json:"template"`
	Field_title       string            `json:"fieldTitle"`
	Field_description string            `json:"fieldDescription"`
	Field_image       string            `json:"fieldImage"`
	Active            bool              `json:"active"`
	Date_Created      time.Time         `json:"date_created"`
	BrokersList       []brokers.Brokers `json:"brokers"`
}
