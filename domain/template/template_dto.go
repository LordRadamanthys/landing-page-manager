package template

type Template struct {
	Id                string `json:"id"`
	Id_mkt            int    `json:"id_mkt"`
	Id_categoria      string `json:"id_categoria"`
	Title             string `json:"title"`
	Template          string `json:"template"`
	Field_title       string `json:"fieldTitle"`
	Field_description string `json:"fieldDescription"`
	Field_image       string `json:"fieldImage"`
	Active            bool   `json:"active"`
	Date_Created      bool   `json:"date_created"`
}
