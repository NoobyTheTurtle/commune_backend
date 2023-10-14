package db

import (
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"gorm.io/gorm"
	"math/rand"
	"time"
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

func mockOffices(offices []*models.Office) {
	rand.Seed(time.Now().UnixNano())
	for i := range offices {
		if rand.Float64() < 0.8 {
			offices[i].Withdrawal = true
		} else {
			offices[i].Withdrawal = false
		}
		if rand.Float64() < 0.8 {
			offices[i].Replenishment = true
		} else {
			offices[i].Replenishment = false
		}
	}
}

func mockAtms(atms []*models.Atm) {
	rand.Seed(time.Now().UnixNano())
	for i := range atms {
		if rand.Float64() < 0.8 {
			atms[i].Withdrawal = true
		} else {
			atms[i].Withdrawal = false
		}
		if rand.Float64() < 0.8 {
			atms[i].Replenishment = true
		} else {
			atms[i].Replenishment = false
		}
	}
}
