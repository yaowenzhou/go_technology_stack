package main

import (
	"fmt"
	hellopb "go_technology_stack/demos/hello_grpc/pb"
	"go_technology_stack/demos/hello_grpc/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "192.168.1.99:9091") // 监听端口
	if nil != err {
		fmt.Println(err.Error())
		return
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)                           // 构建grpc服务端
	hellopb.RegisterHelloServer(grpcServer, &service.HelloServer{}) // 注册Hello服务
	grpcServer.Serve(lis)                                           // 运行grpc服务端提供服务
}
