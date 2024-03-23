package utils

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

const envVarPrefix = ""
const configDelim = "."

type EQConfig struct {
	DatabaseHost     string `koanf:"database.host"`
	DatabaseName     string `koanf:"database.name"`
	DatabaseUser     string `koanf:"database.user"`
	DatabasePassword string `koanf:"database.password"`
	LogLevel         string `koanf:"log.level"`
}

// Global instance of Koanf.
var k = koanf.New(".")
var c *EQConfig

func GetConfig() (*EQConfig, error) {
	if c != nil {
		return c, nil
	}

	var err error

	// Load the configuration
	if err = loadConfig(); err != nil {
		return nil, err
	}

	return c, nil
}

func loadConfig() error {
	err := k.Load(env.Provider(envVarPrefix, configDelim, func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, envVarPrefix)), "_", ".", -1)
	}), nil)

	if err != nil {
		return err
	}

	c = &EQConfig{}

	c.DatabaseHost = k.String("database.host")
	if c.DatabaseHost == "" {

		return fmt.Errorf("database.host is required")
	}

	c.DatabaseName = k.String("database.name")
	if c.DatabaseName == "" {

		return fmt.Errorf("database.name is required")
	}

	c.DatabaseUser = k.String("database.user")
	if c.DatabaseUser == "" {

		return fmt.Errorf("database.user is required")
	}

	c.DatabasePassword = k.String("database.password")
	if c.DatabasePassword == "" {

		return fmt.Errorf("database.password is required")
	}

	return nil
}
