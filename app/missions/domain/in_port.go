package domain

import (
	"context"
	"space-playground/app/astronauts/domain"
	"space-playground/app/shared/infrastructure/graph/model"

	"github.com/google/uuid"
)

type RegisterMissionUseCaseInput struct {
	Title       string
	Description string
	Crew        []domain.Astronaut
}

func (input *RegisterMissionUseCaseInput) FromGql(payload model.NewMissionInput) {
	input = &RegisterMissionUseCaseInput{
		Title:       payload.Title,
		Description: payload.Description,
		Crew:        []domain.Astronaut{},
	}

	for _, crewMemberId := range payload.CrewMembersID {
		input.Crew = append(input.Crew, domain.Astronaut{
			ID: uuid.MustParse(crewMemberId),
		})
	}
}

type RegisterMissionUseCase interface {
	Register(ctx context.Context, input RegisterMissionUseCaseInput) (*int, error)
}
