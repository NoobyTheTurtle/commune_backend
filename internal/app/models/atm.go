package models

import (
	"database/sql/driver"
	"encoding/json"
)

type Atm struct {
	ID            uint     `gorm:"primarykey" json:"id"`
	Address       string   `json:"address"`
	Latitude      float64  `json:"latitude" gorm:"index"`
	Longitude     float64  `json:"longitude" gorm:"index"`
	AllDay        bool     `json:"allDay"`
	Services      Services `gorm:"embedded;embeddedPrefix:services_" json:"services"`
	Withdrawal    bool     `json:"withdrawal"`
	Replenishment bool     `json:"replenishment"`
}

type Services struct {
	Wheelchair        Service `gorm:"type:jsonb" json:"wheelchair"`
	Blind             Service `gorm:"type:jsonb" json:"blind"`
	NfcForBankCards   Service `gorm:"type:jsonb" json:"nfcForBankCards"`
	QrRead            Service `gorm:"type:jsonb" json:"qrRead"`
	SupportsUsd       Service `gorm:"type:jsonb" json:"supportsUsd"`
	SupportsChargeRub Service `gorm:"type:jsonb" json:"supportsChargeRub"`
	SupportsEur       Service `gorm:"type:jsonb" json:"supportsEur"`
	SupportsRub       Service `gorm:"type:jsonb" json:"supportsRub"`
}

type Service struct {
	ServiceCapability string `json:"serviceCapability"`
	ServiceActivity   string `json:"serviceActivity"`
}

func (p *Service) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), p)
}

func (c Service) Value() (driver.Value, error) {
	return json.Marshal(c)
}
