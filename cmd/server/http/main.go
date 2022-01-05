package main

import (
	"flag"
	"fmt"

	"github.com/KonstantinGasser/weeat/api"
	"github.com/KonstantinGasser/weeat/cmd/server/config"
	"github.com/KonstantinGasser/weeat/core/services/records"
	"github.com/KonstantinGasser/weeat/handler"
	"github.com/KonstantinGasser/weeat/pkg/postgres"
	"github.com/sirupsen/logrus"
)

func main() {

	isProd := flag.Bool("prod", false, "will load env from os.ENV")
	// flag.Parse()

	cfg, err := config.Init(".", *isProd)
	if err != nil {
		logrus.Panic(err)
	}
	fmt.Println(cfg)
	// create database dependency
	weeatDB := postgres.New(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, int(cfg.Database.Port))
	if err := weeatDB.Connect(cfg.Database.DatabaseName); err != nil {
		logrus.Panic(err)
	}

	// create API Http server
	apihttp := api.New(
		":8000",
		[]string{"*"},
		[]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
	)

	// create service dependencies
	recordsvc := records.New(nil)

	// setting up routes
	//
	// routes: records.Food
	apihttp.Register("/records/new/food", handler.HandleInsertFood(
		recordsvc,
	), "POST")

	// routes: records.Recipe
	apihttp.Register("/records/new/recipe", handler.HandlerInsertRecipe(
		recordsvc,
	), "POST")

	// start API Http server
	if err := apihttp.Listen(); err != nil {
		logrus.Panic(err)
	}

}
