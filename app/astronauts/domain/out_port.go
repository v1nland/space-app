package domain

import (
	"context"

	"github.com/google/uuid"
)

type CreateAstronautRepository interface {
	Create(ctx context.Context, astronaut *Astronaut) (*uuid.UUID, error)
}

type GetAstronautsRepository interface {
	ById(ctx context.Context, id uuid.UUID) (*Astronaut, error)
	All(ctx context.Context) ([]Astronaut, error)
}
