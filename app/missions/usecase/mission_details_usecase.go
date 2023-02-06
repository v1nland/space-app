package usecase

import (
	"context"
	"space-playground/app/missions/domain"
)

/*
 *	Usecase struct and constructor
 */
type missionDetailsUsecase struct {
	getMissionsRepository domain.ListMissionsUseCase
}

func NewMissionDetailsUsecase(getMissionsRepository domain.GetMissionsRepository) *missionDetailsUsecase {
	return &missionDetailsUsecase{
		getMissionsRepository: getMissionsRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *missionDetailsUsecase) ById(ctx context.Context, id int) (*domain.Mission, error) {
	mission, err := r.getMissionsRepository.ById(ctx, id)
	if err != nil {
		return nil, err
	}

	return mission, nil
}
