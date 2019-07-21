package models

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/local/go-mongo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := os.Getenv("host")
	port := os.Getenv("port")
	connect := fmt.Sprintf("mongodb://%s:%s", host, port)
	conn := options.Client().ApplyURI(connect)
	db, err := mongo.Connect(ctx, conn)
	utils.FindErrors(err, "DB is not connect ")
	DB = db
}
