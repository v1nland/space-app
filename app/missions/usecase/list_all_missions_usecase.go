package usecase

import (
	"context"
	"space-playground/app/missions/domain"
)

/*
 *	Usecase struct and constructor
 */
type listAllMissionsUsecase struct {
	getMissionsRepository domain.GetMissionsRepository
}

func NewListAllMissionsUsecase(getMissionsRepository domain.GetMissionsRepository) *listAllMissionsUsecase {
	return &listAllMissionsUsecase{
		getMissionsRepository: getMissionsRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *listAllMissionsUsecase) All(ctx context.Context) ([]domain.Mission, error) {
	missions, err := r.getMissionsRepository.All(ctx)
	if err != nil {
		return nil, err
	}

	return missions, nil
}
