package utils

import (
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	FindErrors(err, "Cannot Load ENV")
}
