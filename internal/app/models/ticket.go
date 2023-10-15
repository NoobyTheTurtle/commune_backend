package models

import (
	"gorm.io/gorm"
	"math/rand"
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
		Number:   generateTicketNumber(),
		Service:  service,
	}
}

func generateTicketNumber() (ticketNumber string) {
	rand.Seed(time.Now().UnixNano())
	firstChar := 'A' + rune(rand.Intn('Z'-'A'+1))
	secondChar := '0' + rune(rand.Intn('9'-'0'+1))
	thirdChar := '0' + rune(rand.Intn('9'-'0'+1))
	fourthChar := '0' + rune(rand.Intn('9'-'0'+1))
	ticketNumber = string(firstChar) + string(secondChar) + string(thirdChar) + string(fourthChar)
	return
}
