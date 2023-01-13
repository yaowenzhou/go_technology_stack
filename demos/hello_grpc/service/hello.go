package service

import (
	"context"
	hellopb "go_technology_stack/demos/hello_grpc/pb"
)

// HelloServer TODO
type HelloServer struct {
	hellopb.UnimplementedHelloServer
}

// SayHello TODO
func (s *HelloServer) SayHello(
	ctx context.Context, in *hellopb.SayHelloReq,
) (out *hellopb.SayHelloResp, err error) {
	return &hellopb.SayHelloResp{Message: "Hello, Grpc!"}, nil
}
