package main

import (
	"go_technology_stack/demos/hello_go_micro/handler"
	pb "go_technology_stack/demos/hello_go_micro/pb"
	"log"

	"github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Server(grpc.NewServer()),     // 必须使用此语句构造一个grpcserver，否则会概率性返回{Status code: 2 UNKNOWN}
		micro.Address("192.168.1.99:9091"), // 指定微服务的地址
		micro.Name("rpchello.service"),
		micro.Version("latest"),
	)

	pb.RegisterHelloHandler(service.Server(), &handler.Hello{})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
