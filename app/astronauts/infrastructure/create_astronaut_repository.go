package infrastructure

import (
	"space-playground/app/astronauts/domain"
	"space-playground/app/shared/infrastructure/psql/connection"
	"space-playground/app/shared/infrastructure/psql/db_model"

	"github.com/google/uuid"
)

/*
 *	Repository struct and constructor
 */
type createAstronautRepository struct {
	connection connection.PostgreSqlConnection
}

func NewCreateAstronautRepository(connection connection.PostgreSqlConnection) *createAstronautRepository {
	return &createAstronautRepository{
		connection: connection,
	}
}

/*
 *	Repository functions
 */
func (c *createAstronautRepository) Insert(astronaut *domain.Astronaut) (id *uuid.UUID, err error) {
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	db_astronaut := db_model.Astronaut{
		ID:      [16]byte{},
		Name:    astronaut.Name,
		IsPilot: astronaut.IsPilot,
	}

	if err := db.Create(&db_astronaut).Error; err != nil {
		return nil, err
	}

	return &db_astronaut.ID, nil
}
