package brokers

type Brokers struct {
	Id_broker                int    `json:"id_broker"`
	Name                     string `json:"name"`
	Field_broker_title       string `json:"fieldBrokerTitle"`
	Field_broker_description string `json:"fieldBrokerDescription"`
	Field_broker_imageCenter string `json:"fieldBrokerImageCenter"`
	Field_broker_imageLeft   string `json:"fieldBrokerImageLeft"`
	Url_site                 string `json:"site"`
	Url_facebook             string `json:"facebook"`
	Url_instagram            string `json:"instagram"`
	Url_linkedin             string `json:"linkedin"`
	Email                    string `json:"email"`
}
