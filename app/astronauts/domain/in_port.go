package domain

import (
	"context"

	"github.com/google/uuid"
)

type RegisterAstronautUseCase interface {
	Register(ctx context.Context, name string, isPilot bool) (*uuid.UUID, error)
}

type RetrieveAstronautsUseCase interface {
	ById(ctx context.Context, id uuid.UUID) (*Astronaut, error)
}
