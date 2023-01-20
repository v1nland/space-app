package psql

import (
	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/log"

	"space-playground/app/shared/infrastructure/psql/db_model"

	"space-playground/app/shared/infrastructure/psql/connection"
)

func AutoMigrateEntities(conn connection.PostgreSqlConnection) {
	log.Info("autoMigrateEntities...")

	migrate := connection.NewMigrate(conn)
	migrate.AutoMigrateAll(
		db_model.Astronaut{},
		db_model.Mission{},
	)

	log.Info("autoMigrateEntities... OK")
}

func CreatePostgreSqlDbConnection() *connection.PostgreSqlDbConnection {
	host := config.GetString("config.postgresql.host")
	port := config.GetInt("config.postgresql.port")
	database := config.GetString("config.postgresql.database")
	user := config.GetString("config.postgresql.user")
	password := config.GetString("config.postgresql.password")

	connection := connection.NewPostgreSqlConnection(connection.Config().
		Host(host).
		Port(port).
		DatabaseName(database).
		User(user).
		Password(password),
	)
	return connection
}
