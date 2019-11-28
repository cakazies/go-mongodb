package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/local/go-mongo/models"
	"github.com/local/go-mongo/utils"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func init() {
	InitViper()
}

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
		Addr:         ":8000",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}

// InitViper function for initialization package viper
func InitViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	log.Println("Using Config File: ", viper.ConfigFileUsed())
}
