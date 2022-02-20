library gohttp;

import 'dart:convert';
import 'dart:developer';
import 'dart:ffi';
import 'dart:io';

import 'package:call/ffi.dart';
import 'package:ffi/ffi.dart';
import 'package:flutter/foundation.dart';

part 'src/model.dart';
part 'src/godart.dart';
part 'src/gohttp.dart';

/// http 请求对象
class Http {
  // 设置请求根地址
  setBaseAddress(String baseAddr) async {
    goSetBaseAddress(GoString.fromString(baseAddr));
  }

  // 设置超时时间 秒
  setTimeout(int timeout) async {
    goSetTimeout(timeout);
  }

  // 获取gohttp版本信息
  GoHttpVersion getVersion() {
    Pointer<Int8> version = goGetVersion();
    String versionStr = version.cast<Utf8>().toDartString();
    Map<String, dynamic> map = json.decode(versionStr);
    return GoHttpVersion.fromJson(map);
  }

  // post 请求
  post(String url,
      {Map<String, dynamic>? params,
      Map<String, String>? header,
      String? contentType,
      bool? encrypt}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goPost(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
      'content_type': contentType,
      'encrypt': encrypt,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // get 请求
  get(String url,
      {Map<String, dynamic>? params,
      Map<String, String>? header,
      String? contentType,
      bool? encrypt}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goGet(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
      'content_type': contentType,
      'encrypt': encrypt,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // put 请求
  put(String url,
      {Map<String, dynamic>? params,
      Map<String, String>? header,
      String? contentType,
      bool? encrypt}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goPut(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
      'content_type': contentType,
      'encrypt': encrypt,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // delete 请求
  delete(String url,
      {Map<String, dynamic>? params,
      Map<String, String>? header,
      String? contentType,
      bool? encrypt}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goDelete(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
      'content_type': contentType,
      'encrypt': encrypt,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // 转换返回值为ResponseData
  toResponseData(String data) {
    Map<String, dynamic> map = json.decode(data.toString());
    ResponseData responseData = ResponseData.fromJson(map);
    return responseData;
  }
}
