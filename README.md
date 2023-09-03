# todo-be-go


# ■ local setup
## 0. Before setup local
- For using gRPC gateway, needed to add `./proto/google/*` from [googleapis](https://github.com/googleapis/googleapis/tree/master)
- For using validation, needed to add `./proto/buf/validate/*` from [protovalidate-go](https://github.com/bufbuild/protovalidate-go)


## 1. Requirements
- Docker
- Go : `>=v1.20`
  ```
  % go version
  go version go1.20 darwin/arm64
  ```
- Protocol Buffers
  ```
  % brew install protobuf
  % protoc --version
  ```
- install go binaries around Protocol buffers
  ```
  % go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  % go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

  // modules for grpc-gateway
  % go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
  % go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
  % go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  % go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

  // https://github.com/bufbuild/protovalidate-go
  % go install github.com/bufbuild/protovalidate-go@latest
  ```
- set path
  ```
  % vi ~/.zshrc
  --------
  # Add this
  export PATH="$PATH:$(go env GOPATH)/bin"% 
  --------

  % source ~/.zshrc
  ```

- Make go file from proto
  ```
  % make proto-gw
  ```
## 2. 