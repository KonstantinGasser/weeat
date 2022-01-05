package main

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host         string
		Port         int64
		User         string
		Password     string
		DatabaseName string
	}
}

func initConifg(path string, prod bool) (config, error) {
	if prod {
		port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 10)
		if err != nil {
			return config{}, errors.Wrap(err, "os.Getenv of DB_PORT cast error")
		}
		return config{
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
		return config{}, errors.Wrap(err, "viper - read config")
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return config{}, errors.Wrap(err, "viper - unmarshal config")
	}
	return cfg, nil
}
