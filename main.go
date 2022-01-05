package main

import (
	"flag"

	"github.com/KonstantinGasser/weeat/api"
	"github.com/KonstantinGasser/weeat/core/services/records"
	"github.com/KonstantinGasser/weeat/handler"
	"github.com/sirupsen/logrus"
)

func main() {

	hostApi := flag.String("api-host", ":8000", "API host address")
	// hostDB := flag.String("db-host", "localhost", "database hostname")
	// portDB := flag.Int("db-port", 5432, "database port number")
	// dbName := flag.String("db-name", "weeat_db", "name of the database")
	flag.Parse()

	// create database dependency
	// weeatDB := postgres.New("postgres", "weeat_secure", *hostDB, *portDB)
	// if err := weeatDB.Connect(*dbName); err != nil {
	// 	logrus.Panic(err)
	// }

	// create API Http server
	apihttp := api.New(
		*hostApi,
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
