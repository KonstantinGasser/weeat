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
		Port         int64
		User         string
		Password     string
		DatabaseName string
	}
}

func Init(path string, prod bool) (Config, error) {
	if prod {
		port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 10)
		if err != nil {
			return Config{}, errors.Wrap(err, "os.Getenv of DB_PORT cast error")
		}
		return Config{
			Database: struct {
				Host         string
				Port         int64
				User         string
				Password     string
				DatabaseName string
			}{
				Host:         os.Getenv("DB_HOST"),
				Port:         port,
				User:         os.Getenv("DB_USER"),
				Password:     os.Getenv("DB_PASSWORD"),
				DatabaseName: os.Getenv("DB_DATABASE_NAME"),
			},
		}, nil
	}

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
