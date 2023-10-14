package models

import (
	"database/sql/driver"
	"encoding/json"
)

var (
	ServicesList = ServiceList{
		"кредит", "карта", "ипотека", "автокредит", "вклад и счет", "платежи и переводы",
		"страхование", "другие услуги",
	}
)

type Office struct {
	ID                  uint        `gorm:"primarykey" json:"id"`
	SalePointName       string      `json:"salePointName"`
	Address             string      `json:"address"`
	Status              string      `json:"status"`
	OpenHours           HoursList   `gorm:"type:jsonb" json:"openHours"`
	Rko                 string      `json:"rko"`
	OpenHoursIndividual HoursList   `gorm:"type:jsonb" json:"openHoursIndividual"`
	OfficeType          string      `json:"officeType"`
	SalePointFormat     string      `json:"salePointFormat"`
	SuoAvailability     string      `json:"suoAvailability"`
	HasRamp             string      `json:"hasRamp"`
	Latitude            float64     `json:"latitude" gorm:"index"`
	Longitude           float64     `json:"longitude" gorm:"index"`
	MetroStation        string      `json:"metroStation"`
	Distance            int         `json:"distance"`
	Kep                 bool        `json:"kep"`
	MyBranch            bool        `json:"myBranch"`
	Withdrawal          bool        `json:"withdrawal"`
	Replenishment       bool        `json:"replenishment"`
	Services            ServiceList `gorm:"type:jsonb" json:"services"`
}

type ServiceList []string

type HoursList []Hours

type Hours struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

func (hl *ServiceList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), hl)
}

func (hl ServiceList) Value() (driver.Value, error) {
	return json.Marshal(hl)
}

func (hl *HoursList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), hl)
}

func (hl HoursList) Value() (driver.Value, error) {
	return json.Marshal(hl)
}
