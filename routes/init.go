package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/local/go-mongo/models"
	"github.com/spf13/viper"
)

// Route function for routing
func Route() {
	e := echo.New()
	api := e.Group("/api")
	api.POST("/student", models.CreateStudent)
	api.GET("/student", models.GetStudents)
	api.GET("/student/:id", models.GetStudent)
	api.POST("/student/:id/update", models.UpdateStudent)
	api.DELETE("/student/:id", models.DeleteStudent)

	s := &http.Server{
		Addr:         viper.GetString("app.host"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
