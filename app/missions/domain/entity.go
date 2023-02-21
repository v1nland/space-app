package domain

import (
	"space-playground/app/astronauts/domain"
	"space-playground/app/shared/infrastructure/psql/db_model"
)

type Mission struct {
	ID          int
	Title       string
	Description string
	Crew        []domain.Astronaut
}

// add business logic here
func (mission *Mission) FromDbModel(dbMission db_model.Mission) {
	mission.ID = dbMission.ID
	mission.Title = dbMission.Title
	mission.Description = dbMission.Description

	for _, dbCrewMember := range dbMission.Crew {
		mission.Crew = append(mission.Crew, domain.Astronaut{
			ID:      dbCrewMember.ID,
			Name:    dbCrewMember.Name,
			IsPilot: dbCrewMember.IsPilot,
		})
	}
}

func (mission *Mission) ToDbModel() db_model.Mission {
	dbMission := db_model.Mission{
		ID:          mission.ID,
		Title:       mission.Title,
		Description: mission.Description,
	}

	for _, crewMember := range mission.Crew {
		dbMission.Crew = append(dbMission.Crew, db_model.Astronaut{
			ID:      crewMember.ID,
			Name:    crewMember.Name,
			IsPilot: crewMember.IsPilot,
		})
	}

	return dbMission
}
