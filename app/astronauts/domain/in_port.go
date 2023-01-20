package domain

import "github.com/google/uuid"

type RegisterAstronautUseCase interface {
	RegisterAstronaut(name string, isPilot bool) (*uuid.UUID, error)
}
