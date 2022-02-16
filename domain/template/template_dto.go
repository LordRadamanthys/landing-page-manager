package template

type Template struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Template          string `json:"template"`
	Field_title       string `json:"fieldTitle"`
	Field_description string `json:"fieldDescription"`
	Field_image       string `json:"fieldImage"`
}
