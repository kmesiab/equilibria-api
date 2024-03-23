package utils

import (
	"gorm.io/gorm"
)

func PingDatabase(globalDB *gorm.DB) error {

	sqlDB, err := globalDB.DB()

	if err != nil {
		return err
	}

	err = sqlDB.Ping()

	if err != nil {
		return err
	}

	return nil
}
