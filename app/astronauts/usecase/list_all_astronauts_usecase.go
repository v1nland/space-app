package usecase

import (
	"context"
	"space-playground/app/astronauts/domain"
)

/*
 *	Usecase struct and constructor
 */
type listAllAstronautsUsecase struct {
	getAstronautsRepository domain.GetAstronautsRepository
}

func NewListAllAstronautsUsecase(getAstronautsRepository domain.GetAstronautsRepository) *listAllAstronautsUsecase {
	return &listAllAstronautsUsecase{
		getAstronautsRepository: getAstronautsRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *listAllAstronautsUsecase) All(ctx context.Context) ([]domain.Astronaut, error) {
	astronauts, err := r.getAstronautsRepository.All(ctx)
	if err != nil {
		return nil, err
	}

	return astronauts, nil
}
