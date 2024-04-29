package helper

import (
	"log"
)

func Recover() {
	if rec := recover(); rec != nil {
		log.Panicln("recover", rec.(string))
	}
}
