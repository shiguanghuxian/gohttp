LIB_NAME=gohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-s -w -X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"
GOBUILD_DEBUG=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD` -X 'gohttp/gosrc/gohttp.DEBUG=true'"

default:
	@echo 'Usage of make: [ macos | windows | linux | run ]'

run: build
	./${LIB_NAME}

build:
	${GOBUILD_DEBUG} -o ${LIB_NAME} *.go

macos: 
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 ${GOBUILD} -o bin/macos/${LIB_NAME}.dylib -buildmode=c-shared *.go
	cp bin/macos/${LIB_NAME}.dylib ../../example/macos/dylib

ios-arm64: 
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=iphoneos \
	CC=$(PWD)/clangwrap.sh \
	CGO_CFLAGS="-fembed-bitcode" \
	${GOBUILD} -o bin/ios/${LIB_NAME}.dylib -tags ios -buildmode=c-shared *.go
	cp bin/ios/${LIB_NAME}.dylib ../../example/ios/dylib

ios-x86_64:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=amd64 \
	SDK=iphonesimulator \
	CC=$(PWD)/clangwrap.sh \
	${GOBUILD} -o bin/ios/${LIB_NAME}.dylib -tags ios -buildmode=c-shared *.go
	cp bin/ios/${LIB_NAME}.dylib ../../example/ios/dylib

android: 
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 ${GOBUILD} -o bin/android/${LIB_NAME}.so -buildmode=c-shared *.go

linux: 
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 ${GOBUILD} -o bin/linux/${LIB_NAME}.so -buildmode=c-shared *.go

windows: 
	GOOS=windows GOARCH=amd64 ${GOBUILD} -o bin/windows/${LIB_NAME}.dll -buildmode=c-shared .

.PHONY: default macos windows linux run