package brokers

type Brokers struct {
	Id_broker                int    `json:"id_broker"`
	Name                     string `json:"name"`
	Field_broker_title       string `json:"field_broker_title"`
	Field_broker_description string `json:"field_broker_description"`
	Field_broker_imageCenter string `json:"field_broker_imageCenter"`
	Field_broker_imageLeft   string `json:"field_broker_imageLeft"`
	Url_site                 string `json:"site"`
	Url_facebook             string `json:"facebook"`
	Url_instagram            string `json:"instagram"`
	Url_linkedin             string `json:"linkedin"`
	Email                    string `json:"email"`
}
