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
	input.Title = payload.Title
	input.Description = payload.Description

	for _, crewMemberId := range payload.CrewMembersID {
		input.Crew = append(input.Crew, domain.Astronaut{
			ID: uuid.MustParse(crewMemberId),
		})
	}
}

type RegisterMissionUseCase interface {
	Register(ctx context.Context, input RegisterMissionUseCaseInput) (*int, error)
}

type ListMissionsUseCase interface {
	ById(ctx context.Context, id int) (*Mission, error)
}

type ListAllMissionsUseCase interface {
	All(ctx context.Context) ([]Mission, error)
}
