package db

import (
	"commune_backend/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(models.BankBranch{}); err != nil {
		return nil, err
	}

	return db, nil
}
