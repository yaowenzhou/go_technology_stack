package main

import (
	"fmt"
	"go_technology_stack/demos/hello_grpc/pb"
	"go_technology_stack/demos/hello_grpc/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "192.168.1.99:9091")
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServer(grpcServer, &service.HelloServer{})
	grpcServer.Serve(lis)
}
