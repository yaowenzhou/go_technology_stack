hello_grpc:
	go build -o ./demos/bin/hello_grpc ./demos/hello_grpc/main.go

hello_grpc_proto:
	rm -rf ./demos/hello_grpc/pb/
	mkdir ./demos/hello_grpc/pb/
	cd ./demos/hello_grpc/proto/
	powerproto build . -r
