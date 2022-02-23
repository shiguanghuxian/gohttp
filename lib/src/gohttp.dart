part of gohttp;

DynamicLibrary lib = Platform.isMacOS || Platform.isIOS || Platform.isAndroid
    ? DynamicLibrary.open(getGoHttpLibPath())
    : getDyLibModule(getGoHttpLibPath());


String getGoHttpLibPath() {
  if (Platform.isMacOS) {
    return 'libgohttp.dylib';
  }
  if (Platform.isWindows) {
    return 'gosrc/export_c/bin/windows/libgohttp.dll';
  }
  if (Platform.isIOS) {
    return 'libgohttp.dylib';
  }
  if (Platform.isAndroid) {
    return 'libgohttp.so';
  }
  return 'libgohttp.so';
}

// 获取gohttp库版本信息
typedef GetVersionFunc = Pointer<Int8> Function();
GetVersionFunc goGetVersion = lib
    .lookup<NativeFunction<Pointer<Int8> Function()>>('GetVersion')
    .asFunction();

// 设置超时
typedef SetTimeoutFunc = void Function(int);
SetTimeoutFunc goSetTimeout =
    lib.lookup<NativeFunction<Void Function(Int64)>>('SetTimeout').asFunction();

// 设置请求根地址
typedef SetBaseAddressFunc = void Function(Pointer<GoString>);
SetBaseAddressFunc goSetBaseAddress = lib
    .lookup<NativeFunction<Void Function(Pointer<GoString>)>>('SetBaseAddress')
    .asFunction();

// 设置公共请求头
typedef SetHeaderFunc = void Function(Pointer<GoString>);
SetHeaderFunc goSetHeader = lib
    .lookup<NativeFunction<Void Function(Pointer<GoString>)>>('SetHeader')
    .asFunction();

// post 请求
typedef PostFunc = Pointer<Int8> Function(Pointer<GoString>);
PostFunc goPost = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>)>>('Post')
    .asFunction();

// get 请求
typedef GetFunc = Pointer<Int8> Function(Pointer<GoString>);
GetFunc goGet = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>)>>('Get')
    .asFunction();

// put 请求
typedef PutFunc = Pointer<Int8> Function(Pointer<GoString>);
PutFunc goPut = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>)>>('Put')
    .asFunction();

// delete 请求
typedef DeleteFunc = Pointer<Int8> Function(Pointer<GoString>);
DeleteFunc goDelete = lib
    .lookup<NativeFunction<Pointer<Int8> Function(Pointer<GoString>)>>('Delete')
    .asFunction();

// request 请求
typedef RequestFunc = Pointer<Int8> Function(
    Pointer<GoString>, Pointer<GoString>);
RequestFunc goRequest = lib
    .lookup<
        NativeFunction<
            Pointer<Int8> Function(
                Pointer<GoString>, Pointer<GoString>)>>('Request')
    .asFunction();
