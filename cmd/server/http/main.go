package main

import (
	"flag"
	"fmt"

	"github.com/KonstantinGasser/weeat/api"
	"github.com/KonstantinGasser/weeat/cmd/server/config"
	"github.com/KonstantinGasser/weeat/core/services/foodsvc"
	"github.com/KonstantinGasser/weeat/core/services/recipesvc"
	"github.com/KonstantinGasser/weeat/handler"
	"github.com/KonstantinGasser/weeat/pkg/postgres"
	"github.com/sirupsen/logrus"
)

func main() {

	isProd := flag.Bool("prod", false, "will load env from os.ENV")
	flag.Parse()

	// init config from dev(yaml) or prod(ENV)
	var cfg config.Config
	var err error
	if *isProd {
		cfg, err = config.FromEnv()
	} else {
		cfg, err = config.FromYaml(".")
	}
	if err != nil {
		logrus.Panic(err)
	}

	// create database dependency
	weeatDB := postgres.New(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
	if err := weeatDB.Connect(cfg.Database.DatabaseName); err != nil {
		logrus.Panic(err)
	}

	// create API Http server
	apihttp := api.New(
		fmt.Sprintf(":%s", cfg.Server.Port),
		[]string{"*"},
		[]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
	)

	// create service dependencies
	foodsvc := foodsvc.New(weeatDB)
	recipesvc := recipesvc.New(weeatDB)

	// setting up routes
	//
	// routes: food
	apihttp.Register("/api/v1/food", handler.HandleInsertFood(
		foodsvc,
	), "POST")
	apihttp.Register("/api/v1/food/search", handler.HandleSearchFood(
		foodsvc,
	), "GET")
	apihttp.Register("/api/v1/food", handler.HandleGetFood(
		foodsvc,
	), "GET")

	// routes: records.Recipe
	apihttp.Register("/api/v1/recipe", handler.HandlerInsertRecipe(
		recipesvc,
	), "POST")

	// start API Http server
	if err := apihttp.Listen(); err != nil {
		logrus.Panic(err)
	}

}
