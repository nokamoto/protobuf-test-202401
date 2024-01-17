GOBIN=$(shell go env GOPATH)/bin

all: protobuf

$(GOBIN)/buf:
	go install github.com/bufbuild/buf/cmd/buf@v1.28.1

$(GOBIN)/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

protobuf: $(GOBIN)/buf $(GOBIN)/protoc-gen-go
	buf format -w
	buf generate
	go run .
