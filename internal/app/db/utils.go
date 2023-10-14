package db

import (
	"errors"
	"gorm.io/gorm"
)

func IsTableEmpty(db *gorm.DB, models interface{}) (bool, error) {
	result := db.First(models)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, result.Error
	}

	return false, nil
}
