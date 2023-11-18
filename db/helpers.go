package db

import (
	"io/ioutil"
)

func ExecuteSQLFile(filePath string) error {
	sqlQuery, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = GetDB().Exec(string(sqlQuery))
	return err
}

func ReadSQLFile(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}