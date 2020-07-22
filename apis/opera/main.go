package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"opera/handler"
	"opera/client"

	opera "opera/proto/opera"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("fuckopera.api.opera"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Opera service client
		micro.WrapHandler(client.OperaWrapper(service)),
	)

	// Register Handler
	opera.RegisterOperaHandler(service.Server(), new(handler.Opera))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
