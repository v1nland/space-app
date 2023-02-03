package usecase

import (
	"context"
	"space-playground/app/missions/domain"
)

/*
 *	Usecase struct and constructor
 */
type registerMissionUsecase struct {
	createMissionRepository domain.CreateMissionRepository
}

func NewRegisterMissionUsecase(createMissionRepository domain.CreateMissionRepository) *registerMissionUsecase {
	return &registerMissionUsecase{
		createMissionRepository: createMissionRepository,
	}
}

/*
 *	Usecase functions
 */
func (r *registerMissionUsecase) Register(ctx context.Context, input domain.RegisterMissionUseCaseInput) (*int, error) {
	missionEntity := &domain.Mission{
		Title:       input.Title,
		Description: input.Description,
		Crew:        input.Crew,
	}

	missionId, err := r.createMissionRepository.Create(ctx, missionEntity)
	if err != nil {
		return nil, err
	}

	return missionId, nil
}
