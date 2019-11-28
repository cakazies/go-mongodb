package models

import (
	"context"
	"fmt"
	"time"

	"github.com/local/go-mongo/utils"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DB variable connection for mongoDB
	DB *mongo.Client
)

func init() {
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
