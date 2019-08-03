package utils

import (
	"fmt"
	"log"
)

func FailError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func LogError(err error, msg string) {
	if err != nil {
		log.Println(msg, " : ", err)
	}
}

func PanicError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s : %s", msg, err))
	}
}
