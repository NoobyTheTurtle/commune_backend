package filters

import (
	"gorm.io/gorm"
	"strconv"
)

func FilterOfficesByLegalPerson(isLegalPersonStr string) Filter {
	condition, err := strconv.ParseBool(isLegalPersonStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("NOT offices.open_hours @> '[{\"days\": \"Не обслуживает ЮЛ\"}]'")
		} else {
			return db.Where("offices.open_hours @> '[{\"days\": \"Не обслуживает ЮЛ\"}]'")
		}
	}
}

func FilterOfficesByIndividualPerson(isIndividualPersonStr string) Filter {
	condition, err := strconv.ParseBool(isIndividualPersonStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("NOT offices.open_hours_individual @> '[{\"days\": \"Не обслуживает ФЛ\"}]'")
		} else {
			return db.Where("offices.open_hours_individual @> '[{\"days\": \"Не обслуживает ФЛ\"}]'")
		}

	}
}

func FilterOfficesByImmobile(isImmobileStr string) Filter {
	condition, err := strconv.ParseBool(isImmobileStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("offices.has_ramp = 'Y'")
		} else {
			return db.Where("offices.has_ramp = 'N'")
		}

	}
}

func FilterOfficesByPrime(isPrimeStr string) Filter {
	condition, err := strconv.ParseBool(isPrimeStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("offices.office_type LIKE '%Привилегия%'")
		} else {
			return db.Where("offices.office_type NOT LIKE '%Привилегия%'")
		}

	}
}

func FilterOfficesByOpen(isOpenStr, isLegalPersonStr string) Filter {
	isOpen, err := strconv.ParseBool(isOpenStr)
	isLegalPerson, errIsLegalPerson := strconv.ParseBool(isLegalPersonStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isOpen {
			if errIsLegalPerson == nil && isLegalPerson {
				return db.Where("offices.status = 'открытая'")
			} else {
				return db.Where("offices.status = 'открытая'")
			}
		}
		return db.Where("offices.status != 'открытая'")
	}
}

func FilterOfficesByWithdrawal(isWithdrawalStr string) Filter {
	condition, err := strconv.ParseBool(isWithdrawalStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("offices.withdrawal IS TRUE")
		} else {
			return db.Where("offices.withdrawal IS FALSE")
		}
	}
}

func FilterOfficesByReplenishment(isReplenishmentStr string) Filter {
	condition, err := strconv.ParseBool(isReplenishmentStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if condition {
			return db.Where("offices.replenishment IS TRUE")
		} else {
			return db.Where("offices.replenishment IS FALSE")
		}
	}
}
