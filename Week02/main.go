package main

import (
	"log"

	"Week02/service"
)

func main() {
	svr := service.New()
	userinfo, err := svr.GetUserInfo(1)
	if err != nil {
		log.Println("HTTP 500")
		return
	}
	log.Println(userinfo)
}
