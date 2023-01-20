package resolver

import "space-playground/app/astronauts/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	registerAstronautUseCase  domain.RegisterAstronautUseCase
	retrieveAstronautsUsecase domain.RetrieveAstronautsUseCase
}

func NewResolver(
	registerAstronautUseCase domain.RegisterAstronautUseCase,
	retrieveAstronautsUsecase domain.RetrieveAstronautsUseCase,
) *Resolver {
	return &Resolver{
		registerAstronautUseCase:  registerAstronautUseCase,
		retrieveAstronautsUsecase: retrieveAstronautsUsecase,
	}
}
