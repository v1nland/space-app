package connection

import "space-playground/app/shared/infrastructure/log"

type Migrate struct {
	connection Connection
}

func NewMigrate(connection Connection) *Migrate {
	return &Migrate{connection: connection}
}

func (m *Migrate) AutoMigrateAll(tables ...interface{}) {
	db, err := m.connection.GetConnection()
	if err != nil {
		log.WithError(err).Fatal(err.Error())
	}

	err = db.AutoMigrate(tables...)
	if err != nil {
		log.WithError(db.Error).Fatal(db.Error.Error())
	}
}
