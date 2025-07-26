package cmd

import (
	"crud-app/pkg/config"

	"crud-app/pkg/db"
	"log"

	app "crud-app/internal/app"
)

func StartAPIServer() {
	log.Println("Starting API server...")
	config.InitConfig()
	conf := config.GetConfig()
	db.DBConnection(conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Dbname)
	defer db.CloseDBConnection()
	app.StartAPIServer()
}
