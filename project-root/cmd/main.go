package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"github.com/marioTiara/todolistapp/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	uow := repository.NewUnitOfWork(db)
	service := services.NewTaskService(uow)
	handler := handlers.NewHandlers(service)

	e := echo.New()
	v1 := e.Group("v1")

	v1.GET("/", handler.Hello)
	v1.GET("/tasks", handler.GetAllList)
	v1.POST("/tasks", handler.PostTaskHandler)
	v1.GET("/tasks/:id", handler.GetTaskByIDHandler)
	v1.DELETE("/tasks/:id", handler.Delete)
	v1.GET("/subTask/:parentID", handler.GetAllSubListByParentID)

	if err := e.Start(":8080"); err != nil {
		panic("failed to start the server")
	}
}
