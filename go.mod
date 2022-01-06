module github.com/KonstantinGasser/weeat

// +heroku goVersion go1.16
go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx/v4 v4.14.1
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.8.2
	github.com/sirupsen/logrus v1.8.1
	go.mongodb.org/mongo-driver v1.8.1
)

require github.com/spf13/viper v1.10.1
