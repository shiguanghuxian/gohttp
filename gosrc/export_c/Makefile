LIB_NAME=libgohttp
VERSION=`cat '../version'`
GOBUILD=go build -ldflags "-s -w -X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD`"
GOBUILD_DEBUG=go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=$(VERSION) -X 'gohttp/gosrc/gohttp.BUILD_TIME=`date "+%s"`' -X 'gohttp/gosrc/gohttp.GO_VERSION=`go version`' -X gohttp/gosrc/gohttp.GIT_HASH=`git rev-parse HEAD` -X 'gohttp/gosrc/gohttp.DEBUG=true'"
ANDROID_NDK_HOME=${HOME}/Library/Android/sdk/ndk/23.1.7779620

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
	CC=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi21-clang \
	CGO_ENABLED=1 GOOS=android GOARCH=arm ${GOBUILD} -o bin/android/armeabi-v7a/${LIB_NAME}.so -buildmode=c-shared *.go \

	CC=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang \
	CGO_ENABLED=1 GOOS=android GOARCH=arm64 ${GOBUILD} -o bin/android/arm64-v8a/${LIB_NAME}.so -buildmode=c-shared *.go \

	CC=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android21-clang \
	CGO_ENABLED=1 GOOS=android GOARCH=amd64 ${GOBUILD} -o bin/android/x86_64/${LIB_NAME}.so -buildmode=c-shared *.go \

	CC=${ANDROID_NDK_HOME}/toolchains/llvm/prebuilt/darwin-x86_64/bin/i686-linux-android21-clang \
	CGO_ENABLED=1 GOOS=android GOARCH=386 ${GOBUILD} -o bin/android/x86/${LIB_NAME}.so -buildmode=c-shared *.go
	
	cp bin/android/x86_64/${LIB_NAME}.so ../../example/android/app/src/main/jniLibs/x86_64
	cp bin/android/x86/${LIB_NAME}.so ../../example/android/app/src/main/jniLibs/x86
	cp bin/android/arm64-v8a/${LIB_NAME}.so ../../example/android/app/src/main/jniLibs/arm64-v8a
	cp bin/android/armeabi-v7a/${LIB_NAME}.so ../../example/android/app/src/main/jniLibs/armeabi-v7a
	

linux: 
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 ${GOBUILD} -o bin/linux/${LIB_NAME}.so -buildmode=c-shared *.go

windows: 
	GOOS=windows GOARCH=amd64 ${GOBUILD} -o bin/windows/${LIB_NAME}.dll -buildmode=c-shared *.go

.PHONY: default macos windows linux run