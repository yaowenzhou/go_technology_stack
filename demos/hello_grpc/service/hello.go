package service

import (
	"context"
	"go_technology_stack/demos/hello_grpc/pb"
)

// HelloServer TODO
type HelloServer struct {
	pb.UnimplementedHelloServer
}

// SayHello TODO
func (s *HelloServer) SayHello(ctx context.Context, in *pb.SayHelloReq) (out *pb.SayHelloResp, err error) {
	return &pb.SayHelloResp{Message: "Hello, Grpc!"}, nil
}
