# gRPC Example Project

This project demonstrates a basic gRPC setup in Go, including a simple `SayHello` service.

## Project Structure

```
grpc-example/
├── go.mod
├── proto/
│   └── hello/v1/hello.proto
├── server.go
└── client.go
```

## Prerequisites

- Go 1.22 or later
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Installation

### Install Go

Run the following commands to install Go:

```sh
brew install go
```

### Install Protobuf Binaries

Run the following commands to install the Protobuf compiler:

```sh
brew install protobuf
```

### Install Go Protocol Buffers Plugins

Run the following commands to install the necessary Go code-generation plugins:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
```

### Install gRPC Debugging Tools

Run the following commands to install the necessary plugins:

```sh
brew install grpcurl
brew install grpcui
```

Make sure the `$GOPATH/bin` directory is in your `PATH` so that the `protoc` compiler can find the plugins.

### Generate Go Code from `.proto` Files

Navigate to the project root directory and run the following command:

```sh
protoc --go_out=. --go-grpc_out=. proto/hello/v1/hello.proto
```

## Running the Server

Navigate to the project root directory and run the following command:

```sh
go run server.go
```

## Running the Client

In a separate terminal, navigate to the project root directory and run the following command:

```sh
go run client.go
```