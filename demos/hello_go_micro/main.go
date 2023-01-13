package main

import (
	"go_technology_stack/demos/hello_go_micro/handler"
	pb "go_technology_stack/demos/hello_go_micro/pb"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
)

func main() {
	rpcServer := server.NewServer(
		server.Name("rpchello.service"),
		server.Address("0.0.0.0:8001"),
	)
	pb.RegisterHelloHandler(rpcServer, &handler.Hello{})
	service := micro.NewService(
		micro.Server(rpcServer),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
