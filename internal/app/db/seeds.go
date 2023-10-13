package db

import (
	"commune_backend/internal/app/models"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

type InitOptions struct {
	FilePath string
	Targets  interface{}
}

func initSeeds(db *gorm.DB) error {
	var offices []*models.Office
	var atms []models.Atm

	//if err := mockOffices(); err != nil {
	//	return err
	//}
	//
	//if err := mockAtms(); err != nil {
	//	return err
	//}

	if err := initModel(db, "seeds/offices.json", &offices); err != nil {
		return err
	}

	if err := initModel(db, "seeds/atms.json", &atms); err != nil {
		return err
	}

	return nil
}

func initModel(db *gorm.DB, filepath string, targets interface{}) error {
	empty, err := isDatabaseEmpty(db, targets)
	if err != nil {
		return err
	}

	if !empty {
		return nil
	}

	if err = readJSONFile(filepath, targets); err != nil {
		return err
	}

	db.Create(targets)

	return nil
}

func readJSONFile(filepath string, targets interface{}) error {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteValue, targets); err != nil {
		return err
	}

	return nil
}

func isDatabaseEmpty(db *gorm.DB, targets interface{}) (bool, error) {
	result := db.First(targets)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, result.Error
	}

	return false, nil
}

//func mockOffices() error {
//	var offices []*models.Office
//	err := readJSONFile("seeds/offices.json", &offices)
//	if err != nil {
//		return err
//	}
//
//	rand.Seed(time.Now().UnixNano())
//	for i := range offices {
//		offices[i].TypeClients = "all"
//		if rand.Float64() <= 0.7 {
//			offices[i].LimitedMobility = true
//		} else {
//			offices[i].LimitedMobility = false
//		}
//
//		if rand.Float64() <= 0.5 {
//			offices[i].TypeClients = "individuals"
//		}
//
//		if rand.Float64() <= 0.2 {
//			offices[i].TypeClients = "entities"
//		}
//
//		if rand.Float64() <= 0.4 {
//			offices[i].PremiumService = true
//		} else {
//			offices[i].PremiumService = false
//		}
//	}
//
//	jsonData, err := json.MarshalIndent(offices, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	err = os.WriteFile("seeds/mock_offices.json", jsonData, 0644)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func mockAtms() error {
//	var atms []*models.Atm
//	err := readJSONFile("seeds/atms.json", &atms)
//	if err != nil {
//		return err
//	}
//	rand.Seed(time.Now().UnixNano())
//	for i := range atms {
//		if rand.Float64() <= 0.5 {
//			atms[i].LimitedMobility = true
//		} else {
//			atms[i].LimitedMobility = false
//		}
//	}
//
//	jsonData, err := json.MarshalIndent(atms, "", "  ")
//	if err != nil {
//		return err
//	}
//
//	err = os.WriteFile("seeds/mock_atms.json", jsonData, 0644)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
