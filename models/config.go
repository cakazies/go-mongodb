package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/local/go-mongo/utils"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB      *mongo.Client
	cfgFile string
)

func init() {
	InitViper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	connect := fmt.Sprintf("mongodb://%s:%s", host, port)
	conn := options.Client().ApplyURI(connect)
	db, err := mongo.Connect(ctx, conn)
	utils.PanicError(err, "DB is not connect ")
	DB = db
}

func InitViper() {
	viper.SetConfigFile("toml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./../configs")
		viper.SetConfigName("config")
	}
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	utils.FailError(err, "Error Viper config")
	log.Println("Using Config File: ", viper.ConfigFileUsed())
}
