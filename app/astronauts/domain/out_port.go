package domain

import "github.com/google/uuid"

type CreateAstronautRepository interface {
	Insert(astronaut *Astronaut) (id *uuid.UUID, err error)
}
