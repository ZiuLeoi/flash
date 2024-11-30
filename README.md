# flash
A distributed framework based on golang

## Generate RPC Code
```shell
protoc --go_out=. --go-grpc_out=. proto/task.proto
```