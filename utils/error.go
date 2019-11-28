package utils

import (
	"fmt"
	"log"
)

// FailError function for setting fail error
func FailError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

// LogError function for log error
func LogError(err error, msg string) {
	if err != nil {
		log.Println(msg, " : ", err)
	}
}

// PanicError function for panic error
func PanicError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s : %s", msg, err))
	}
}
