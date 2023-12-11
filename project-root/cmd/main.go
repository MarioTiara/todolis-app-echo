package main

import (
	"fmt"

	"github.com/bxcodec/faker"
	"github.com/marioTiara/todolistapp/config"
	"github.com/marioTiara/todolistapp/internal/api/models"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"github.com/marioTiara/todolistapp/internal/platform/database"
	"github.com/marioTiara/todolistapp/internal/platform/server"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

func main() {
	task := generateFakeTask(3)
	fmt.Println(task)
	configuration, _ := config.LoadConfig("../config")
	postgress := database.NewPostGressDB(configuration)
	store := storages.NewLocalStoarge("uploads")
	uow := repository.NewUnitOfWork(postgress.GetDB())
	service := services.NewServices(uow, store)

	server, _ := server.NewServer(configuration, service)
	server.Start()

}

func generateFakeTask(depth int) models.Task {
	var task models.Task

	if depth <= 0 {
		return task
	}

	// Generate fake data for the Task struct
	err := faker.FakeData(&task)
	if err != nil {
		fmt.Println("Error generating fake Task data:", err)
	}

	// Manually customize or exclude specific fields if needed
	task.Title = "CustomTitle"
	task.Description = "CustomDescription"

	// Generate fake data for the Files struct
	for i := 0; i < 3; i++ {
		var file models.Files
		err := faker.FakeData(&file)
		if err != nil {
			fmt.Println("Error generating fake Files data:", err)
		}

		// Append the generated file to the task's Files slice
		task.Files = append(task.Files, file)
	}

	// Generate fake data for children with reduced depth
	for i := 0; i < 2; i++ {
		childTask := generateFakeTask(depth - 1)
		if childTask.ID != 0 {
			task.Children = append(task.Children, childTask)
		}
	}

	return task
}
