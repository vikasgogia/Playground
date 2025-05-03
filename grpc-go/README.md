# Welcome to the gRPC-go tutorial

## Protobuf (Protocol Buffer)

Protobuf is a language-agnostic, compact binary format developed by Google for serializing structured data.

- Used to define gRPC messages and services via .proto files.
- Much faster and smaller than JSON or XML.
- Comes with a compiler (protoc) that generates strongly typed code (e.g., Go structs).

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
