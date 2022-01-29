LIB_NAME=gohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"

default:
	@echo 'Usage of make: [ build | debug ]'

build: 
	GOOS=js GOARCH=wasm ${GOBUILD}  *.go

run: 
	go run *.go

.PHONY: default build debug