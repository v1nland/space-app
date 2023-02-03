package psql

import (
	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/log"

	"space-playground/app/shared/infrastructure/psql/db_model"

	"space-playground/app/shared/infrastructure/psql/connection"
)

func AutoMigrateEntities(conn connection.Connection) {
	log.Info("autoMigrateEntities...")

	migrate := connection.NewMigrate(conn)
	migrate.AutoMigrateAll(
		db_model.Astronaut{},
		db_model.Mission{},
	)

	log.Info("autoMigrateEntities... OK")
}

func CreateDbConnection() *connection.DbConnection {
	host := config.Values.Postgres.Host
	port := config.Values.Postgres.Port
	database := config.Values.Postgres.Database
	user := config.Values.Postgres.User
	password := config.Values.Postgres.Password

	connection := connection.NewConnection(
		connection.Config().
			Host(host).
			Port(port).
			DatabaseName(database).
			User(user).
			Password(password),
	)

	return connection
}
