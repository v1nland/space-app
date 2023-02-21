package main

import (
	"net/http"
	astronautsInfrastructure "space-playground/app/astronauts/infrastructure"
	astronautsUsecase "space-playground/app/astronauts/usecase"
	missionsInfrastructure "space-playground/app/missions/infrastructure"
	missionsUsecase "space-playground/app/missions/usecase"
	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/graph"
	"space-playground/app/shared/infrastructure/graph/resolver"
	"space-playground/app/shared/infrastructure/log"
	"space-playground/app/shared/infrastructure/psql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func init() {
	config.LoadSettings()
}

func main() {
	// create logger
	log.Init()
	log.Info("%s is starting", config.Values.App.Name)

	// create psql connection
	psqlConnection := psql.CreateDbConnection()
	psql.AutoMigrateEntities(psqlConnection)

	// create repositories
	createAstronautRepository := astronautsInfrastructure.NewCreateAstronautRepository(psqlConnection)
	getAstronautsRepository := astronautsInfrastructure.NewGetAstronautsRepository(psqlConnection)
	createMissionRepository := missionsInfrastructure.NewCreateAstronautRepository(psqlConnection)
	getMissionsRepository := missionsInfrastructure.NewGetMissionsRepository(psqlConnection)

	// create usecases
	registerAstronautUseCase := astronautsUsecase.NewRegisterAstronautUsecase(createAstronautRepository)
	astronautDetailsUseCase := astronautsUsecase.NewAstronautDetailsUsecase(getAstronautsRepository)
	listAllAstronautsUseCase := astronautsUsecase.NewListAllAstronautsUsecase(getAstronautsRepository)
	registerMissionUseCase := missionsUsecase.NewRegisterMissionUsecase(createMissionRepository)
	missionDetailsUseCase := missionsUsecase.NewMissionDetailsUsecase(getMissionsRepository)

	// create graphql resolver
	resolver := resolver.NewResolver(
		registerAstronautUseCase,
		astronautDetailsUseCase,
		listAllAstronautsUseCase,
		registerMissionUseCase,
		missionDetailsUseCase,
	)

	// create graphql server
	server := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: resolver,
			},
		),
	)

	// init server
	http.Handle("/", playground.Handler("space-playground", "/graphql"))
	http.Handle("/graphql", server)

	log.Info("[%s] is ready to handle messages and listen in port: 8080", config.Values.App.Name)
	log.Fatal(http.ListenAndServe(":8080", nil).Error())
}
