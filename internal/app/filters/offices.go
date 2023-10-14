package filters

import (
	"gorm.io/gorm"
	"strconv"
)

func FilterOfficesByLegalPerson(isLegalPersonStr string) Filter {
	isLegalPerson, err := strconv.ParseBool(isLegalPersonStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isLegalPerson {
			return db.Where("NOT offices.open_hours @> '[{\"days\": \"Не обслуживает ЮЛ\"}]'")
		} else {
			return db.Where("offices.open_hours @> '[{\"days\": \"Не обслуживает ЮЛ\"}]'")
		}
	}
}

func FilterOfficesByIndividualPerson(isIndividualPersonStr string) Filter {
	isIndividualPerson, err := strconv.ParseBool(isIndividualPersonStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isIndividualPerson {
			return db.Where("NOT offices.open_hours_individual @> '[{\"days\": \"Не обслуживает ФЛ\"}]'")
		} else {
			return db.Where("offices.open_hours_individual @> '[{\"days\": \"Не обслуживает ФЛ\"}]'")
		}

	}
}

func FilterOfficesByImmobile(isImmobileStr string) Filter {
	isImmobile, err := strconv.ParseBool(isImmobileStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isImmobile {
			return db.Where("offices.has_ramp = 'Y'")
		} else {
			return db.Where("offices.has_ramp = 'N'")
		}

	}
}

func FilterOfficesByPrime(isPrimeStr string) Filter {
	isPrime, err := strconv.ParseBool(isPrimeStr)
	return func(db *gorm.DB) *gorm.DB {
		if err != nil {
			return db
		}
		if isPrime {
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
