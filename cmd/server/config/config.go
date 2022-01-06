package config

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host         string
		Port         int
		User         string
		Password     string
		DatabaseName string
	}
	Server struct {
		Port string
	}
}

func FromEnv() (Config, error) {
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, errors.Wrap(err, "os.Getenv of DB_PORT cast error")
	}
	return Config{
		Database: struct {
			Host         string
			Port         int
			User         string
			Password     string
			DatabaseName string
		}{
			Host:         os.Getenv("DB_HOST"),
			Port:         dbport,
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
		},
		Server: struct {
			Port string
		}{
			Port: os.Getenv("PORT"),
		},
	}, nil
}

func FromYaml(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".weeat")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, "viper - read Config")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, errors.Wrap(err, "viper - unmarshal Config")
	}
	return cfg, nil
}
