package models

type Weather struct {
	ID          int    `json:"id"`
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Conditions  string `json:"conditions"`
}
