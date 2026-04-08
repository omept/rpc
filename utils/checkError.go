package utils

import "log"

func CheckFatalError(e error) {
	if e != nil {
		log.Fatalln("Fatal error: ", e.Error())
	}
}
