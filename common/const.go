package common

import "log"

const (
	DbTypeNote       = 1
	DbTypeUser       = 2
	DbTypeRestaurant = 3
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover Err", err)
	}
}
