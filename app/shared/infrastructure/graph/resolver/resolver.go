package resolver

import "space-playground/app/astronauts/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	registerAstronautUseCase domain.RegisterAstronautUseCase
}

func NewResolver(registerAstronautUseCase domain.RegisterAstronautUseCase) *Resolver {
	return &Resolver{
		registerAstronautUseCase: registerAstronautUseCase,
	}
}
