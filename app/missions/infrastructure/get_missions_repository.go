package infrastructure

import (
	"context"
	astronautsDomain "space-playground/app/astronauts/domain"
	missionsDomain "space-playground/app/missions/domain"
	"space-playground/app/shared/infrastructure/psql/connection"
	"space-playground/app/shared/infrastructure/psql/db_model"

	"gorm.io/gorm/clause"
)

/*
 *	Repository struct and constructor
 */
type getMissionsRepository struct {
	connection connection.Connection
}

func NewGetMissionsRepository(connection connection.Connection) *getMissionsRepository {
	return &getMissionsRepository{
		connection: connection,
	}
}

/*
 *	Repository functions
 */
func (c *getMissionsRepository) ById(ctx context.Context, id int) (mission *missionsDomain.Mission, err error) {
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	db_mission := db_model.Mission{}

	if err := db.Preload(clause.Associations).Where("id = ?", id).First(&db_mission).Error; err != nil {
		return nil, err
	}

	mission = &missionsDomain.Mission{
		ID:          db_mission.ID,
		Title:       db_mission.Title,
		Description: db_mission.Description,
		Crew:        []astronautsDomain.Astronaut{},
	}

	for _, crewMember := range db_mission.Crew {
		mission.Crew = append(mission.Crew, astronautsDomain.Astronaut{
			ID:      crewMember.ID,
			Name:    crewMember.Name,
			IsPilot: crewMember.IsPilot,
		})
	}

	return mission, nil
}
