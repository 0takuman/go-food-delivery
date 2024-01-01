package common

import "log"

const (
	DbTypeNote = 1
	DbTypeUser = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover Err", err)
	}
}
