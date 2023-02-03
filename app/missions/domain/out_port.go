package domain

import "context"

type CreateMissionRepository interface {
	Create(ctx context.Context, mission *Mission) (id *int, err error)
}
