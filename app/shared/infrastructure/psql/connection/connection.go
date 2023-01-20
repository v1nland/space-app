package connection

import (
	"errors"

	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const LogMode = 1

var connection *gorm.DB

type PostgreSqlConnection interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
}

type PostgreSqlDbConnection struct {
	opts *PostgreSqlOptions
	url  string
}

func NewPostgreSqlConnection(opts ...*PostgreSqlOptions) *PostgreSqlDbConnection {
	databaseOptions := MergeOptions(opts...)
	url := databaseOptions.GetUrlConnection()

	if url == "" {
		log.Fatal(errors.New("error creating connection, empty url").Error())
	}

	return &PostgreSqlDbConnection{
		opts: databaseOptions,
		url:  url,
	}
}

func (r *PostgreSqlDbConnection) GetConnection() (*gorm.DB, error) {
	var err error

	if connection == nil || !isAlive() {
		log.Info("trying to connect to database")

		if connection, err = gorm.Open(postgres.Open(r.url), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: config.GetString("config.gorm.prefix"), // table name prefix, table for `User` would be `t_users`
			},
		}); err != nil {
			log.WithError(err).Error("error trying to connect to DB")
			return nil, err
		} else {
			log.Info("connected to database")
		}
	}

	connection.Logger.LogMode(LogMode)
	return connection, nil
}

func (r *PostgreSqlDbConnection) CloseConnection() {

	if sqlDb, err := connection.DB(); err != nil {
		if err = sqlDb.Close(); err != nil {
			log.WithError(err).Error("error trying to close connection")
		}
	} else {
		log.Info("connection closed")
	}
}

func isAlive() bool {
	if sqlDb, err := connection.DB(); err != nil {
		if err = sqlDb.Ping(); err != nil {
			log.WithError(err).Error("error trying to ping to database")
			return false
		}
	}
	return true
}
