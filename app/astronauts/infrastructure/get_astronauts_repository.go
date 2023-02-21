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
func (c *getAstronautsRepository) ById(ctx context.Context, id uuid.UUID) (*domain.Astronaut, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// retrieve astronaut
	db_astronaut := db_model.Astronaut{}
	if err := db.Where("id = ?", id).First(&db_astronaut).Error; err != nil {
		return nil, err
	}

	// convert to domain model
	astronaut := domain.Astronaut{}
	astronaut.FromDbModel(db_astronaut)

	// return
	return &astronaut, nil
}

func (c *getAstronautsRepository) All(ctx context.Context) ([]domain.Astronaut, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// retrieve astronauts
	db_astronauts := []db_model.Astronaut{}
	if err := db.Find(&db_astronauts).Error; err != nil {
		return nil, err
	}

	// convert to domain model
	astronauts := []domain.Astronaut{}

	for _, db_astronaut := range db_astronauts {
		astronaut := domain.Astronaut{}
		astronaut.FromDbModel(db_astronaut)

		astronauts = append(astronauts, astronaut)
	}

	// return
	return astronauts, nil
}
