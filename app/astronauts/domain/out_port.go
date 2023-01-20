package domain

import (
	"context"

	"github.com/google/uuid"
)

type CreateAstronautRepository interface {
	Create(ctx context.Context, astronaut *Astronaut) (id *uuid.UUID, err error)
}

type GetAstronautsRepository interface {
	ById(ctx context.Context, id uuid.UUID) (*Astronaut, error)
}
