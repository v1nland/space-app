package usecase

import (
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
 *	Repository functions
 */
func (r *registerAstronautUsecase) Insert(name string, isPilot bool) (*uuid.UUID, error) {
	astronaut := domain.Astronaut{
		Name:    name,
		IsPilot: isPilot,
	}

	id, err := r.createAstronautRepository.Insert(&astronaut)
	if err != nil {
		return nil, err
	}

	return id, nil
}
