package usecase

import (
	"context"
	"space-playground/app/astronauts/domain"

	"github.com/google/uuid"
)

/*
 *	Usecase struct and constructor
 */
type retrieveAstronautsUsecase struct {
	getAstronautsRepository domain.GetAstronautsRepository
}

func NewRetrieveAstronautsUsecase(getAstronautsRepository domain.GetAstronautsRepository) *retrieveAstronautsUsecase {
	return &retrieveAstronautsUsecase{
		getAstronautsRepository: getAstronautsRepository,
	}
}

/*
 *	Repository functions
 */
func (r *retrieveAstronautsUsecase) ById(ctx context.Context, id uuid.UUID) (*domain.Astronaut, error) {
	astronaut, err := r.getAstronautsRepository.ById(ctx, id)
	if err != nil {
		return nil, err
	}

	return astronaut, nil
}
