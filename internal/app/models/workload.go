package models

import "time"

type Workload struct {
	ID         uint      `gorm:"primarykey" json:"-"`
	DateTime   time.Time `json:"dateTime"`
	OfficeId   uint      `json:"officeId"`
	Status     bool      `json:"status"`
	Clients    uint      `json:"clients"`
	Employers  uint      `json:"employers"`
	MaxClients uint      `json:"maxClients"`
	Percent    float32   `json:"percent"`
}
