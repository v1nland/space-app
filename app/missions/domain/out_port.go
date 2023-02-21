package domain

import "context"

type CreateMissionRepository interface {
	Create(ctx context.Context, mission *Mission) (*int, error)
}

type GetMissionsRepository interface {
	ById(ctx context.Context, id int) (*Mission, error)
	All(ctx context.Context) ([]Mission, error)
}
