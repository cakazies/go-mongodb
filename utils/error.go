package utils

import "log"

func FindErrors(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
