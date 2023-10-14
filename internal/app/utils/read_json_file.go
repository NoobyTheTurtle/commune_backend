package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONFile(filepath string, targets interface{}) error {
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
