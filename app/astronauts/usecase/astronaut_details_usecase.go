package usecase

import (
	"context"
	"space-playground/app/astronauts/domain"

	"github.com/google/uuid"
)

/*
 *	Usecase struct and constructor
 */
type astronautDetailsUsecase struct {
	getAstronautsRepository domain.GetAstronautsRepository
}

func NewAstronautDetailsUsecase(getAstronautsRepository domain.GetAstronautsRepository) *astronautDetailsUsecase {
	return &astronautDetailsUsecase{
		getAstronautsRepository: getAstronautsRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *astronautDetailsUsecase) ById(ctx context.Context, id uuid.UUID) (*domain.Astronaut, error) {
	astronaut, err := r.getAstronautsRepository.ById(ctx, id)
	if err != nil {
		return nil, err
	}

	return astronaut, nil
}
