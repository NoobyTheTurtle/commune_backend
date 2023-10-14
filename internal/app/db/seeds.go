package db

import (
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"gorm.io/gorm"
)

func initSeeds(db *gorm.DB) error {
	var offices []*models.Office
	var atms []*models.Atm

	if err := initOffices(db, "seeds/offices.json", &offices); err != nil {
		return err
	}
	if err := initAtms(db, "seeds/atms.json", &atms); err != nil {
		return err
	}

	return nil
}

func initOffices(db *gorm.DB, filepath string, offices *[]*models.Office) error {
	empty, err := IsTableEmpty(db, offices)
	if err != nil {
		return err
	}

	if !empty {
		return nil
	}

	if err = utils.ReadJSONFile(filepath, offices); err != nil {
		return err
	}

	mockOffices(*offices)

	db.Create(offices)

	return nil
}

func initAtms(db *gorm.DB, filepath string, atms *[]*models.Atm) error {
	empty, err := IsTableEmpty(db, atms)
	if err != nil {
		return err
	}

	if !empty {
		return nil
	}

	if err = utils.ReadJSONFile(filepath, atms); err != nil {
		return err
	}

	mockAtms(*atms)

	db.Create(atms)

	return nil
}
