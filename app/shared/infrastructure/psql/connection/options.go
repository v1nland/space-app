package connection

import (
	"fmt"

	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/log"
)

type PostgreSqlOptions struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *PostgreSqlOptions {
	return new(PostgreSqlOptions)
}

func (c *PostgreSqlOptions) DatabaseName(databaseName string) *PostgreSqlOptions {
	c.databaseName = &databaseName
	return c
}

func (c *PostgreSqlOptions) Host(host string) *PostgreSqlOptions {
	c.host = &host
	return c
}

func (c *PostgreSqlOptions) Port(port int) *PostgreSqlOptions {
	c.port = &port
	return c
}

func (c *PostgreSqlOptions) User(user string) *PostgreSqlOptions {
	c.user = &user
	return c
}

func (c *PostgreSqlOptions) Password(password string) *PostgreSqlOptions {
	c.password = &password
	return c
}

func MergeOptions(opts ...*PostgreSqlOptions) *PostgreSqlOptions {
	option := new(PostgreSqlOptions)

	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}
		if opt.host != nil {
			option.host = opt.host
		}
		if opt.port != nil {
			option.port = opt.port
		}
		if opt.user != nil {
			option.user = opt.user
		}
		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

var (
	defaultPort = 5432
)

func (d *PostgreSqlOptions) GetUrlConnection() string {
	urlFormat := "postgresql://%v:%v@%v:%v/%v"

	if d.port == nil {
		d.port = &defaultPort
	}

	if config.Values.Postgres.Ssl == false {
		urlFormat = "postgresql://%v:%v@%v:%v/%v?sslmode=disable"
	}

	log.Info("connection: %s", fmt.Sprintf(urlFormat, *d.user, "************", *d.host, *d.port, *d.databaseName))
	return fmt.Sprintf(urlFormat, *d.user, *d.password, *d.host, *d.port, *d.databaseName)
}
