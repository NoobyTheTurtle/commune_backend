package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	OfficeId int64 `json:"officeId"`
	UserId   int64 `json:"userId"`
}
