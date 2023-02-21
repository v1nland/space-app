package domain

import (
	"space-playground/app/shared/infrastructure/psql/db_model"

	"github.com/google/uuid"
)

type Astronaut struct {
	ID      uuid.UUID
	Name    string
	IsPilot bool
}

// add business logic here
func (astronaut *Astronaut) FromDbModel(dbAstronaut db_model.Astronaut) {
	astronaut.ID = dbAstronaut.ID
	astronaut.Name = dbAstronaut.Name
	astronaut.IsPilot = dbAstronaut.IsPilot
}

func (astronaut *Astronaut) ToDbModel() db_model.Astronaut {
	return db_model.Astronaut{
		ID:      astronaut.ID,
		Name:    astronaut.Name,
		IsPilot: astronaut.IsPilot,
	}
}
