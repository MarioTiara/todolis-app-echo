package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
	"github.com/marioTiara/todolistapp/internal/api/services"
)

func SetRoutes(e *echo.Echo, s services.Service) {
	handler := handlers.NewHandlers(s)
	v1 := e.Group("v1")
	v1.GET("/", handler.Accessible)
	v1.POST("/login", handler.Login)
	v1.GET("/tasks", handler.GetAllList)
	v1.GET("/tasks/:id", handler.GetTaskByIDHandler)
	v1.GET("/subTask/:parentID", handler.GetAllSubListByParentID)
	v1.GET("/uploads/get", handler.DownloadFile)

	v1.POST("/tasks", handler.PostTaskHandler)
	v1.POST("/uploads/add", handler.UploadTaskFilesHandler)
	v1.PUT("/tasks/:id", handler.Update)

	v1.DELETE("/tasks/:id", handler.DeleteTask)
	v1.DELETE("/uploads/delete/:id", handler.DeleteFile)
}
