LIB_NAME=gohttp
VERSION=`cat 'version'`
GOBUILD=go build -ldflags "-X main.VERSION=$(VERSION) -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X main.GIT_HASH=`git rev-parse HEAD`"

default:
	@echo 'Usage of make: [ macos | windows | linux | run | clean ]'

run:
	go run *.go

macos: 
	go build -o bin/macos/${LIB_NAME}.dylib -buildmode=c-shared *.go
	cp bin/macos/${LIB_NAME}.dylib ../example/macos/dylib
linux: 
	go build -o bin/linux/${LIB_NAME}.so -buildmode=c-shared *.go

windows: 
	go build -o bin/windows/${LIB_NAME}.dll -buildmode=c-shared .

clean: 
	@rm -f ./bin/${LIB_NAME}*

.PHONY: default macos windows linux run clean