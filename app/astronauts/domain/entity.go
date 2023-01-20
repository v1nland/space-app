package domain

import "github.com/google/uuid"

type Astronaut struct {
	ID      uuid.UUID
	Name    string
	IsPilot bool
}
