hello_grpc:
	go build ./demos/hello_grpc/main.go
hello_grpc_proto:
	cd ./demos/hello_grpc/proto/
	powerproto build . -r