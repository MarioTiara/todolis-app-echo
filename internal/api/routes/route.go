package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/marioTiara/todolistapp/internal/api/handlers"
)

func SetRoutes(g *echo.Group, handler *handlers.Handler) {

	g.GET("/tasks", handler.GetAllList)
	g.GET("/tasks/:id", handler.GetTaskByIDHandler)
	g.GET("/subTask/:parentID", handler.GetAllSubListByParentID)
	g.GET("/uploads/get", handler.DownloadFile)

	g.POST("/tasks", handler.PostTaskHandler)
	g.POST("/uploads/add", handler.UploadTaskFilesHandler)
	g.PUT("/tasks", handler.Update)

	g.DELETE("/tasks/:id", handler.DeleteTask)
	g.DELETE("/uploads/delete/:id", handler.DeleteFile)
}
