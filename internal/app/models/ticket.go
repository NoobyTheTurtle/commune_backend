package models

import (
	"commune_backend/internal/app/utils"
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	OfficeId  int64          `json:"officeId"`
	UserId    int64          `json:"userId"`
	Number    string         `json:"number"`
	Service   string         `json:"service"`
}

func NewTicket(officeId, userId int64, service string) *Ticket {
	return &Ticket{
		OfficeId: officeId,
		UserId:   userId,
		Number:   utils.GenerateTicketNumber(),
		Service:  service,
	}
}
