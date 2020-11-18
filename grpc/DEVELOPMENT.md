# Development

Health monitoring by sending requests to an endpoint on the application.

## Generate gRPC service

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:grpc/ \
  -I grpc grpc/proto/health/health.proto
```

## Order

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:grpc/ \
  -I grpc grpc/proto/order/order.proto
```