package main

import (
	"net/http"
	"space-playground/app/astronauts/infrastructure"
	"space-playground/app/astronauts/usecase"
	"space-playground/app/shared/infrastructure/config"
	"space-playground/app/shared/infrastructure/graph"
	"space-playground/app/shared/infrastructure/graph/resolver"
	"space-playground/app/shared/infrastructure/log"
	"space-playground/app/shared/infrastructure/psql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func init() {
	config.LoadSettings("SmartSide", "Monumental")
}

func main() {
	// create logger
	log.Init()
	log.Info("%s is starting", config.GetString("app"))

	// create psql connection
	psqlConnection := psql.CreatePostgreSqlDbConnection()
	psql.AutoMigrateEntities(psqlConnection)

	// create repositories
	createAstronautRepository := infrastructure.NewCreateAstronautRepository(psqlConnection)

	// create usecases
	registerAstronautUseCase := usecase.NewRegisterAstronautUsecase(createAstronautRepository)

	// create graphql resolver
	resolver := resolver.NewResolver(registerAstronautUseCase)

	// create graphql server
	server := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: resolver,
			},
		),
	)

	// init server
	http.Handle("/", playground.Handler("Playground", "/graphql"))
	http.Handle("/graphql", server)

	http.ListenAndServe(":8080", nil)
}
