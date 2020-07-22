package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"opera/handler"
	"opera/subscriber"

	opera "opera/proto/opera"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("fuckopera.service.opera"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	opera.RegisterOperaHandler(service.Server(), new(handler.Opera))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("fuckopera.service.opera", service.Server(), new(subscriber.Opera))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
