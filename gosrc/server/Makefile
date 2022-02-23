VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"

default:
	@echo 'Usage of make: [ build | debug ]'

build: 
	${GOBUILD} -o server-demo *.go

run: build
	./server-demo

.PHONY: default build debug