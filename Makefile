UNAME := $(shell uname)
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	gopkg.in/alecthomas/gometalinter.v2 \
	github.com/golang/protobuf/proto \
	github.com/golang/protobuf/ptypes/struct \
	google.golang.org/grpc \
	github.com/gogo/protobuf/proto \
	github.com/gogo/protobuf/jsonpb \
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/gogo/protobuf/gogoproto \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

### /home/gheis/goApps/src/github.com/gallactic/hubble-service

PACKAGES=$(shell go list ./... | grep -v '/vendor/')
INCLUDE = -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com -I=${GOPATH}/src/github.com/gallactic/hubble_service
INCLUDE2 = -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gallactic/hubble_service -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOPATH = --proto_path=${GOPATH}/src:${GOPATH}/src/github.com/gogo/protobuf/protobuf:.
HUBBLE = ${GOPATH}/src/github.com/gallactic/hubble_service

### Tools & dependencies
tools:
	go get $(GOTOOLS)
	@gometalinter.v2 --install

########################################
### Protobuf
proto:

	--protoc $(PROTOPATH) --gogo_out=plugins=grpc:. $(HUBBLE)/blockchain.proto

########################################
### Formatting, linting, and vetting
fmt:
	@go fmt ./...