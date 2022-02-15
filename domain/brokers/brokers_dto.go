package brokers

type Brokers struct {
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	Field_broker_title       string `json:"fieldBrokerTitle"`
	Field_broker_description string `json:"fieldBrokerDescription"`
	Field_broker_imageCenter string `json:"fieldBrokerImageCenter"`
	Field_broker_imageLeft   string `json:"fieldBrokerImageLeft"`
}
