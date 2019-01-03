UNAME := $(shell uname)
GOTOOLS = \
	github.com/golang/dep/cmd/dep \
	gopkg.in/alecthomas/gometalinter.v2 \
	google.golang.org/grpc \
	github.com/golang/protobuf/proto \
	github.com/gogo/protobuf/gogoproto

### /home/gheis/goApps/src/github.com/gallactic/hubble-service

PACKAGES=$(shell go list ./... | grep -v '/vendor/')
INCLUDE = -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com -I=${GOPATH}/src/github.com/gallactic/hubble_service

### Tools & dependencies
tools:
	go get $(GOTOOLS)
	@gometalinter.v2 --install

########################################
### Protobuf
proto:

	--protoc $(INCLUDE) --gogo_out=plugins=grpc:. ./blockchain.proto

########################################
### Formatting, linting, and vetting
fmt:
	@go fmt ./...