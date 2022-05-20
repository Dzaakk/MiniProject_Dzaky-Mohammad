package models

import "gorm.io/gorm"

type Restaurants struct {
	gorm.Model
	RestName string `json: "restaurant name" form: "restaurant name"`
	Cuisine  string `json: "cuisine" form: "cuisine"`
	Address  string `json: "address" form: "address"`
	Contact  string `json: "contact" form: "contact"`
}
