package db

import (
	"commune_backend/internal/app/models"
	"commune_backend/internal/app/utils"
	"gorm.io/gorm"
)

func initSeeds(db *gorm.DB) error {
	var (
		offices   []*models.Office
		atms      []*models.Atm
		workloads []*models.Workload
	)

	if err := initOffices(db, "seeds/offices.json", &offices); err != nil {
		return err
	}
	if err := initAtms(db, "seeds/atms.json", &atms); err != nil {
		return err
	}
	if err := initWorkloads(db, "seeds/workloads.json", &workloads); err != nil {
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

	if result := db.Create(offices); result.Error != nil {
		return result.Error
	}

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

	if result := db.Create(atms); result.Error != nil {
		return result.Error
	}

	return nil
}

func initWorkloads(db *gorm.DB, filepath string, workloads *[]*models.Workload) error {
	empty, err := IsTableEmpty(db, workloads)
	if err != nil {
		return err
	}

	if !empty {
		return nil
	}

	if err = utils.ReadJSONFile(filepath, workloads); err != nil {
		return err
	}

	size := 100
	var j int
	for i := 0; i < len(*workloads); i += size {
		j += size
		if j > len(*workloads) {
			j = len(*workloads)
		}
		if result := db.Create((*workloads)[i:j]); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
