package filters

import (
	"gorm.io/gorm"
	"strconv"
)

func FilterAtmsByImmobile(isImmobileStr string) Filter {
	isImmobile, err := strconv.ParseBool(isImmobileStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isImmobile {
			return db.Where("atms.services_wheelchair = '{\"serviceActivity\": \"AVAILABLE\", \"serviceCapability\": \"SUPPORTED\"}'")
		} else {
			return db.Where("atms.services_wheelchair = '{\"serviceActivity\": \"UNAVAILABLE\", \"serviceCapability\": \"UNSUPPORTED\"}'")
		}
	}
}
