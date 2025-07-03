package dependencies

import (
	"golang-crud-api/internals/infra"
	"golang-crud-api/internals/interfaces/handlers"
	"golang-crud-api/internals/repositories"
	"golang-crud-api/internals/usecases"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

func Setup() *dig.Container {
	container := dig.New()

	if err := container.Provide(infra.NewMongoDatabase); err != nil {
		log.Fatalf("error while registering Mongodb\n\tError: %v", err)
	}

	if err := container.Provide(func(db *mongo.Database) repositories.TaskRepository {
		return repositories.NewTaskRepository(db)
	}); err != nil {
		log.Fatalf("error while registering TaskRepository\n\tError: %v", err)
	}

	if err := container.Provide(usecases.NewTaskUseCase); err != nil {
		log.Fatalf("error while registering UserUseCase\n\tError: %v", err)
	}
	if err := container.Provide(handlers.NewTaskHandler); err != nil {
		log.Fatalf("error while registering UserHandler\n\tError: %v", err)
	}
	return container
}
