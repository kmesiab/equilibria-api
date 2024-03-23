package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	klogger "github.com/kmesiab/go-klogger"
)

var globalDB *gorm.DB

// InitDB initializes the database connection using the provided configuration.
func InitDB(config *EQConfig) (*gorm.DB, error) {

	// Format the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabaseName,
	)

	logLevel := klogger.StringToLogLevel(config.LogLevel)

	// Open the database with the MySQL driver
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
}

// GetDB returns the database connection instance.
func GetDB(config *EQConfig) *gorm.DB {
	var err error

	if globalDB == nil {

		globalDB, err = InitDB(config)

		// If we don't have a database connection, panic
		if err != nil {
			panic(fmt.Sprintf("failed to initialize database: %s\n", err))
		}

		// If we can't reach the database, panic
		if PingDatabase(globalDB) != nil {
			msg := fmt.Sprintf("failed to ping database: %s\n", err)
			panic(msg)
		} else {
			klogger.Logf("Database connected").Info()
		}
	}

	return globalDB
}
