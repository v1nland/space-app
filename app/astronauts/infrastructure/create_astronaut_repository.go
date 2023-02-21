package infrastructure

import (
	"context"
	"space-playground/app/astronauts/domain"
	"space-playground/app/shared/infrastructure/psql/connection"

	"github.com/google/uuid"
)

/*
 *	Repository struct and constructor
 */
type createAstronautRepository struct {
	connection connection.Connection
}

func NewCreateAstronautRepository(connection connection.Connection) *createAstronautRepository {
	return &createAstronautRepository{
		connection: connection,
	}
}

/*
 *	Repository functions
 */
func (c *createAstronautRepository) Create(ctx context.Context, astronaut *domain.Astronaut) (*uuid.UUID, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// convert to db model
	db_astronaut := astronaut.ToDbModel()
	if err := db.Create(&db_astronaut).Error; err != nil {
		return nil, err
	}

	// return
	return &db_astronaut.ID, nil
}
