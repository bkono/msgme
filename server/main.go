package main

import (
	"github.com/bkono/msgme/proto/msgme"
	"github.com/bkono/msgme/server/handler"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.msgme"),
		micro.Version("latest"),
	)

	// Register Handler
	msgme.RegisterMsgMeHandler(service.Server(), new(handler.MsgMe))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
