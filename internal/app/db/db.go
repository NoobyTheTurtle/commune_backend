package db

import (
	"commune_backend/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(url string, logger logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		models.Office{},
		models.Atm{},
		models.Ticket{},
		models.Workload{},
	); err != nil {
		return nil, err
	}

	if err = initSeeds(db); err != nil {
		return nil, err
	}

	return db, nil
}
