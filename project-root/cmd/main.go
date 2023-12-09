package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/config"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/internal/api/services"
	"github.com/marioTiara/todolistapp/internal/platform/database"
	"github.com/marioTiara/todolistapp/internal/platform/storages"
	"github.com/marioTiara/todolistapp/internal/repository"
)

func main() {
	// dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	configuration, _ := config.LoadConfig("./config")
	postgress := database.NewPostGressDB(configuration.DbSource)
	store := storages.NewLocalStoarge("uploads")
	uow := repository.NewUnitOfWork(postgress.GetDB())
	service := services.NewTaskService(uow, store)
	handler := handlers.NewHandlers(service)

	e := echo.New()
	v1 := e.Group("v1")
	e.POST("/upload", handleUpload)
	v1.GET("/", handler.Hello)
	v1.GET("/tasks", handler.GetAllList)
	v1.POST("/tasks", handler.PostTaskHandler)
	v1.GET("/tasks/:id", handler.GetTaskByIDHandler)
	v1.PUT("/tasks/:id", handler.Update)
	v1.DELETE("/tasks/:id", handler.Delete)
	v1.GET("/subTask/:parentID", handler.GetAllSubListByParentID)

	if err := e.Start(":8080"); err != nil {
		panic("failed to start the server")
	}
}

type RequestData struct {
	Name string `json:"name"`
}

func handleUpload(c echo.Context) error {
	// Parse JSON body
	var requestData RequestData
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON body"})
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileName, _ := storages.NewLocalStoarge("uploads").SaveFile(file)

	return c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", fileName))
}
