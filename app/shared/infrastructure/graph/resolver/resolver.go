package resolver

import (
	astronautsDomain "space-playground/app/astronauts/domain"
	missionsDomain "space-playground/app/missions/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	registerAstronautUseCase astronautsDomain.RegisterAstronautUseCase
	astronautDetailsUsecase  astronautsDomain.GetAstronautByIdUseCase
	listAllAstronautsUseCase astronautsDomain.ListAllAstronautsUseCase
	registerMissionUseCase   missionsDomain.RegisterMissionUseCase
	missionDetailsUseCase    missionsDomain.ListMissionsUseCase
}

func NewResolver(
	registerAstronautUseCase astronautsDomain.RegisterAstronautUseCase,
	astronautDetailsUsecase astronautsDomain.GetAstronautByIdUseCase,
	listAllAstronautsUseCase astronautsDomain.ListAllAstronautsUseCase,
	registerMissionUseCase missionsDomain.RegisterMissionUseCase,
	missionDetailsUseCase missionsDomain.ListMissionsUseCase,
) *Resolver {
	return &Resolver{
		registerAstronautUseCase: registerAstronautUseCase,
		astronautDetailsUsecase:  astronautDetailsUsecase,
		listAllAstronautsUseCase: listAllAstronautsUseCase,
		registerMissionUseCase:   registerMissionUseCase,
		missionDetailsUseCase:    missionDetailsUseCase,
	}
}
