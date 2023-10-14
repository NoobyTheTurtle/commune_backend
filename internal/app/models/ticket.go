package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	OfficeId  uint           `json:"-"`
	Office    Office         `json:"office"`
	UserId    uint           `json:"userId"`
}
