package main

import (
	proto "github.com/Rakanixu/factorial-micro-service/server/proto"
	"github.com/Rakanixu/factorial-micro-service/server/handler"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.factorial"),
		micro.Flags(
			cli.StringFlag{
				Name: "factorial",
				Usage: "The factorial",
			},
		),
	)

	service.Init()

	proto.RegisterFactorialHandler(service.Server(), new(handler.Factorial))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
