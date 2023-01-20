package main

import (
	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/log"
	"space-playground/app/shared/infrastructure/psql"
)

const defaultPort = "8080"

func init() {
	config.LoadSettings("SmartSide", "Monumental")
}

func main() {
	log.Init()
	log.Info("%s is starting", config.GetString("app"))

	// create psql connection
	psqlConnection := psql.CreatePostgreSqlDbConnection()
	psql.AutoMigrateEntities(psqlConnection)
}
