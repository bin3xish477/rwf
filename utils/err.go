package utils

import "log"

func Must(err error) {
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
