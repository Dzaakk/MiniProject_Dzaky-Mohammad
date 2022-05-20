package models

import "gorm.io/gorm"

type Reviews struct {
	gorm.Model
	UserID  int    `json: "userid" form:"userid"`
	RestID  int    `json: "restid" form: "restid"`
	Comment string `json: "comment" form: "comment"`
	Rating  int    `json: "rating" form: "rating"`
}
