package filters

import (
	"gorm.io/gorm"
	"strconv"
)

func FilterAtmsByImmobile(isImmobileStr string) Filter {
	condition, err := strconv.ParseBool(isImmobileStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("atms.services_wheelchair = '{\"serviceActivity\": \"AVAILABLE\", \"serviceCapability\": \"SUPPORTED\"}'")
		} else {
			return db.Where("atms.services_wheelchair = '{\"serviceActivity\": \"UNAVAILABLE\", \"serviceCapability\": \"UNSUPPORTED\"}'")
		}
	}
}

func FilterAtmsByWithdrawal(isWithdrawalStr string) Filter {
	condition, err := strconv.ParseBool(isWithdrawalStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("atms.withdrawal IS TRUE")
		} else {
			return db.Where("atms.withdrawal IS FALSE")
		}
	}
}

func FilterAtmsByReplenishment(isReplenishmentStr string) Filter {
	condition, err := strconv.ParseBool(isReplenishmentStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("atms.replenishment IS TRUE")
		} else {
			return db.Where("atms.replenishment IS FALSE")
		}
	}
}
