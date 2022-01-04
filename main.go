package main

import (
	"flag"

	"github.com/KonstantinGasser/lowco/api"
	"github.com/KonstantinGasser/lowco/core/services/records"
	"github.com/KonstantinGasser/lowco/handler"
	"github.com/sirupsen/logrus"
)

func main() {

	hostApi := flag.String("api-host", ":8000", "API host address")
	// hostDB := flag.String("db-host", "localhost", "database hostname")
	// portDB := flag.Int("db-port", 5432, "database port number")
	// dbName := flag.String("db-name", "lowco_db", "name of the database")
	flag.Parse()

	// create database dependency
	// lowcoDB := postgres.New("postgres", "lowco_secure", *hostDB, *portDB)
	// if err := lowcoDB.Connect(*dbName); err != nil {
	// 	logrus.Panic(err)
	// }

	// create API Http server
	apihttp := api.New(
		*hostApi,
		[]string{"*"},
		[]string{"OPTIONS", "GET", "POST", "PUT"},
	)

	// create service dependencies
	recordsvc := records.New(nil)

	// setting up routes
	//
	// routes: records.Food
	apihttp.Register("/records/food", handler.HandleInsertFood(
		recordsvc,
	), "PUT")
	// apihttp.Register("/records/food", nil)
	// apihttp.Register("/records/food", nil)
	// apihttp.Register("/records/food", nil)
	// apihttp.Register("records/search", nil)

	// routes: records.Recipe
	apihttp.Register("/records/recipe", handler.HandlerInsertRecipe(
		recordsvc,
	), "PUT")
	// apihttp.Register("/records/recipe", nil)
	// apihttp.Register("/records/recipe", nil)
	// apihttp.Register("/records/recipe", nil)

	// start API Http server
	if err := apihttp.Listen(); err != nil {
		logrus.Panic(err)
	}

}
