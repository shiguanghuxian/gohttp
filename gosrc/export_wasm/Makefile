LIB_NAME=libgohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-s -w -X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"
GOBUILD_DEBUG=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD` -X 'gohttp/gosrc/gohttp.DEBUG=true'"

default:
	@echo 'Usage of make: [ build | debug ]'

# 减小体积 brotli -o gohttp.wasm.br gohttp.wasm
build: 
	GOOS=js GOARCH=wasm ${GOBUILD} -o ../server/static/${LIB_NAME}.wasm *.go
	gzip -9 < ../server/static/${LIB_NAME}.wasm | wc -c

# 日志打印
debug: 
	GOOS=js GOARCH=wasm ${GOBUILD_DEBUG} -o ../server/static/${LIB_NAME}.wasm *.go
	gzip -9 < ../server/static/${LIB_NAME}.wasm | wc -c

.PHONY: default build debug