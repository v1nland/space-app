package infrastructure

import (
	"context"
	"space-playground/app/astronauts/domain"
	"space-playground/app/shared/infrastructure/psql/connection"
	"space-playground/app/shared/infrastructure/psql/db_model"

	"github.com/google/uuid"
)

/*
 *	Repository struct and constructor
 */
type getAstronautsRepository struct {
	connection connection.Connection
}

func NewGetAstronautsRepository(connection connection.Connection) *getAstronautsRepository {
	return &getAstronautsRepository{
		connection: connection,
	}
}

/*
 *	Repository functions
 */
func (c *getAstronautsRepository) ById(ctx context.Context, id uuid.UUID) (astronaut *domain.Astronaut, err error) {
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	db_astronaut := db_model.Astronaut{}

	if err := db.Where("id = ?", id).First(&db_astronaut).Error; err != nil {
		return nil, err
	}

	astronaut = &domain.Astronaut{
		ID:      db_astronaut.ID,
		Name:    db_astronaut.Name,
		IsPilot: db_astronaut.IsPilot,
	}

	return astronaut, nil
}
