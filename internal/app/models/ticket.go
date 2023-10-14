package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	OfficeId  int64          `json:"officeId"`
	UserId    int64          `json:"userId"`
}
