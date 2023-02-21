package infrastructure

import (
	"context"
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
func (c *getMissionsRepository) ById(ctx context.Context, id int) (*missionsDomain.Mission, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// retrieve mission
	db_mission := db_model.Mission{}
	if err := db.Preload(clause.Associations).Where("id = ?", id).First(&db_mission).Error; err != nil {
		return nil, err
	}

	// convert to domain model
	mission := missionsDomain.Mission{}
	mission.FromDbModel(db_mission)

	// return
	return &mission, nil
}

func (c *getMissionsRepository) All(ctx context.Context) ([]missionsDomain.Mission, error) {
	// get connection
	db, err := c.connection.GetConnection()
	if err != nil {
		return nil, err
	}

	// retrieve missions
	db_missions := []db_model.Mission{}
	if err := db.Preload(clause.Associations).Find(&db_missions).Error; err != nil {
		return nil, err
	}

	// convert to domain model
	missions := []missionsDomain.Mission{}

	for _, db_mission := range db_missions {
		mission := missionsDomain.Mission{}
		mission.FromDbModel(db_mission)

		missions = append(missions, mission)
	}

	// return
	return missions, nil
}
