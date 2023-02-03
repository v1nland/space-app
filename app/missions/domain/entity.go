package domain

import "space-playground/app/astronauts/domain"

type Mission struct {
	ID          int
	Title       string
	Description string
	Crew        []domain.Astronaut
}
