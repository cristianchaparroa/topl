package models

type Search struct {
	Term string	`json:"term"`
	Lat string	`json:"lat"`
	Lng string	`json:"lng"`
}
