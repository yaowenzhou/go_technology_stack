package handler

import (
	"context"
	hellopb "go_technology_stack/demos/hello_go_micro/pb"
)

type Hello struct{}

func (s *Hello) SayHello(ctx context.Context,
	in *hellopb.SayHelloReq, out *hellopb.SayHelloResp,
) error {
	out.Message = "Hello go-micro"
	out.Ext = "I'm yaowenzhou"
	return nil
}
