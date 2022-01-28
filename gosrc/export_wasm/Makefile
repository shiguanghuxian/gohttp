LIB_NAME=gohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"

default:
	@echo 'Usage of make: [ build ]'

build: 
	GOOS=js GOARCH=wasm ${GOBUILD} -o server/static/${LIB_NAME}.wasm wasm/*.go

run:
	go run server/main.go

.PHONY: default build