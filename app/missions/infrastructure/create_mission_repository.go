package infrastructure

import (
	"context"
	"space-playground/app/missions/domain"
	"space-playground/app/shared/infrastructure/psql/connection"
)

/*
 *	Repository struct and constructor
 */
type createMissionRepository struct {
	connection connection.Connection
}

func NewCreateAstronautRepository(connection connection.Connection) *createMissionRepository {
	return &createMissionRepository{
		connection: connection,
	}
}

/*
 *	Repository functions
 */
func (c *createMissionRepository) Create(ctx context.Context, mission *domain.Mission) (*int, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// convert to db model
	db_mission := mission.ToDbModel()
	if err := db.Create(&db_mission).Error; err != nil {
		return nil, err
	}

	// return
	return &db_mission.ID, nil
}
