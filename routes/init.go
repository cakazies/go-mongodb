package routes

import (
	"github.com/labstack/echo"
	"github.com/local/go-mongo/models"
)

func Route() {
	e := echo.New()
	api := e.Group("/api")
	api.POST("/student", models.CreateStudent)
	api.GET("/student", models.GetStudents)
	api.GET("/student/:id", models.GetStudent)
	api.POST("/student/:id/update", models.UpdateStudent)
	api.DELETE("/student/:id", models.DeleteStudent)
	e.Start(":8000")
}
