package filters

import "gorm.io/gorm"

type Filter func(*gorm.DB) *gorm.DB

type Filters []Filter

func (fs *Filters) Scopes() []func(*gorm.DB) *gorm.DB {
	var scopes []func(*gorm.DB) *gorm.DB
	for _, f := range *fs {
		scopes = append(scopes, f)
	}
	return scopes
}
