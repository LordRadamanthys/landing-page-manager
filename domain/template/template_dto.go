package template

type Template struct {
	Id                string `json:"id" bson:"_id,omitempty"`
	Id_mkt            int    `json:"id_mkt"`
	Id_category       string `json:"id_category"`
	Title             string `json:"title"`
	Template          string `json:"template"`
	Field_title       string `json:"fieldTitle"`
	Field_description string `json:"fieldDescription"`
	Field_image       string `json:"fieldImage"`
	Active            bool   `json:"active"`
	Date_Created      string `json:"date_created"`
}
