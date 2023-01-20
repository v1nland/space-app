package domain

import "space-playground/app/astronauts/domain"

type Mission struct {
	ID          int64
	Title       string
	Description string
	Crew        []domain.Astronaut
}
