package cmd

import (
	"crud-app/pkg/config"

	"crud-app/pkg/db"
	"log"

	app "crud-app/internal/app"

	"github.com/spf13/cobra"
)

var apiCmd = cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	Long:  `This command starts the API server for the CRUD application.`,
	Run:   startAPIServer,
}

func init() {
	RootCmd.AddCommand(&apiCmd)
}

func startAPIServer(cmd *cobra.Command, args []string) {
	log.Println("Starting API server...")
	config.InitConfig()
	conf := config.GetConfig()
	db.DBConnection(conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.Database)
	defer db.CloseDBConnection()
	app.StartAPIServer()
}
