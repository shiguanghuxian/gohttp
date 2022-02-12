LIB_NAME=gohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"

default:
	@echo 'Usage of make: [ macos | windows | linux | run ]'

run: build
	./${LIB_NAME}

build:
	${GOBUILD} -o ${LIB_NAME} *.go

macos: 
	GOOS=darwin GOARCH=amd64 ${GOBUILD} -o bin/macos/${LIB_NAME}.dylib -buildmode=c-shared *.go
	cp bin/macos/${LIB_NAME}.dylib ../../example/macos/dylib

linux: 
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o bin/linux/${LIB_NAME}.so -buildmode=c-shared *.go

windows: 
	GOOS=windows GOARCH=amd64 ${GOBUILD} -o bin/windows/${LIB_NAME}.dll -buildmode=c-shared .

.PHONY: default macos windows linux run