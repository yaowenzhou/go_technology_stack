hello_grpc:
	go build -o ./demos/bin/hello_grpc ./demos/hello_grpc/main.go

hello_grpc_proto:
	rm -rf ./demos/hello_grpc/pb/
	mkdir ./demos/hello_grpc/pb/
	cd ./demos/hello_grpc/proto/
	powerproto build . -r

hello_go_micro:
	go build -o ./demos/bin/hello_go_micro ./demos/hello_go_micro/main.go

hello_go_micro_proto:
	rm -rf ./demos/hello_go_micro/pb/
	mkdir ./demos/hello_go_micro/pb/
	cd ./demos/hello_go_micro/proto/
	powerproto build . -r