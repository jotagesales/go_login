/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jotagesales/pkg/database"
	"github.com/jotagesales/pkg/models"
	"github.com/jotagesales/pkg/routes"
	"github.com/jotagesales/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var port string

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: not use hard corded for this information, refactor to config
		db, err := database.Connect("postgres", "mysecretpassword", "go_login")

		if err != nil {
			log.Fatalf("could not connect database detail: %s", err)
		}

		admin := models.User{Name: "admin", Email: "test@login.com", Password: "mysecretpassword"}
		db.AutoMigrate(&models.User{})
		db.Create(&admin)

		// this affect load test
		engine := gin.New()
		// engine := gin.Default()
		route := routes.GetRoutes(engine, db)

		gin.SetMode(gin.ReleaseMode)

		s := server.NewServer(route, port)
		server.Runserver(s)
		defer db.Close()
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runserverCmd.Flags().StringVarP(&port, "port", "p", ":8080", "port to run api")
}
