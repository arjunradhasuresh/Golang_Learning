package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use:   "Crud",
	Short: "A simple CRUD application",
	Long:  `This application provides basic CRUD operations for managing resources.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
