# Welcome to the gRPC-go tutorial

## Install Go (mac)

```bash
$ brew install go

$ go version

$ go env
```

## Install protobuf

```bash
$ brew install protobuf

# protobuf compiler
$ protoc --version
```

## Install gRPC go plugins

```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Update PATH

```bash
# appends go bin to PATH system variable
$ export PATH="$PATH:$(go env GOPATH)/bin"
```
