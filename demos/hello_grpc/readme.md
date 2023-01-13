# Hello Grpc!

## 1. 安装两个插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 2. 目录结构如下
```bash
hello_grpc/
├── main.go
├── pb
│   ├── hello_grpc.pb.go
│   └── hello.pb.go
├── proto
│   ├── hello.proto
│   └── powerproto.yaml
├── readme.md
└── service
    └── hello.go
```

## 3. 编写proto文件
文件内容请参考[./proto/hello.proto](./proto/hello.proto)

## 4. 执行命令编译proto
执行如下命令，相关参数的具体含义请自行百度，后续如果有时间，我会补齐

```bash
cd hello_grpc/proto
protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative hello.proto
```

## 5. 在service中实现Hello接口
请参考 [./service/hello.go](./service/hello.go)

## 6. 编写main.go
请参考 [./main.go](./main.go)