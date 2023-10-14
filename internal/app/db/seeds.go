package db

import (
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"errors"
	"gorm.io/gorm"
)

func initSeeds(db *gorm.DB) error {
	var offices []*models.Office
	var atms []*models.Atm

	if err := initModel(db, "seeds/offices.json", &offices); err != nil {
		return err
	}
	if err := initModel(db, "seeds/atms.json", &atms); err != nil {
		return err
	}

	return nil
}

func initModel(db *gorm.DB, filepath string, models interface{}) error {
	empty, err := isTableEmpty(db, models)
	if err != nil {
		return err
	}

	if !empty {
		return nil
	}

	if err = utils.ReadJSONFile(filepath, models); err != nil {
		return err
	}

	db.Create(models)

	return nil
}

func isTableEmpty(db *gorm.DB, models interface{}) (bool, error) {
	result := db.First(models)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, result.Error
	}

	return false, nil
}
