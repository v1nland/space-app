package usecase

import (
	"context"
	"space-playground/app/astronauts/domain"

	"github.com/google/uuid"
)

/*
 *	Usecase struct and constructor
 */
type registerAstronautUsecase struct {
	createAstronautRepository domain.CreateAstronautRepository
}

func NewRegisterAstronautUsecase(createAstronautRepository domain.CreateAstronautRepository) *registerAstronautUsecase {
	return &registerAstronautUsecase{
		createAstronautRepository: createAstronautRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *registerAstronautUsecase) Register(ctx context.Context, name string, isPilot bool) (*uuid.UUID, error) {
	astronaut := domain.Astronaut{
		Name:    name,
		IsPilot: isPilot,
	}

	id, err := r.createAstronautRepository.Create(ctx, &astronaut)
	if err != nil {
		return nil, err
	}

	return id, nil
}
