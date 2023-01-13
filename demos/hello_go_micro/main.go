package main

import (
	"go_technology_stack/demos/hello_go_micro/handler"
	pb "go_technology_stack/demos/hello_go_micro/pb"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/server"
)

func main() {
	rpcServer := server.NewServer(
		server.Name("rpchello.service"),
		server.Address("192.168.1.99:8080"),
	)
	pb.RegisterHelloHandler(rpcServer, &handler.Hello{})
	service := micro.NewService(
		micro.Server(rpcServer),
		micro.Broker(broker.NewBroker(broker.Addrs("192.168.1.99:8081"))),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
