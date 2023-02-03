package infrastructure

import (
	"context"
	"space-playground/app/missions/domain"
	"space-playground/app/shared/infrastructure/psql/connection"
	"space-playground/app/shared/infrastructure/psql/db_model"
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
func (c *createMissionRepository) Create(ctx context.Context, mission *domain.Mission) (id *int, err error) {
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	db_mission := db_model.Mission{
		Title:       mission.Title,
		Description: mission.Description,
		Crew:        []db_model.Astronaut{},
	}

	for _, crewMember := range mission.Crew {
		db_mission.Crew = append(db_mission.Crew, db_model.Astronaut{
			ID: crewMember.ID,
		})
	}

	if err := db.Create(&db_mission).Error; err != nil {
		return nil, err
	}

	return &db_mission.ID, nil
}
