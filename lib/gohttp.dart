library gohttp;

import 'dart:convert';
import 'dart:developer';
import 'dart:ffi';

import 'package:ffi/ffi.dart';
import 'package:flutter/foundation.dart';
import 'package:gohttp/src/godart.dart';
import 'package:gohttp/src/gohttp.dart';

part 'src/model.dart';

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

  // post 请求
  post(String url, {Map<String, dynamic>? params, Map<String, String>? header}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goPost(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // get 请求
  get(String url, {Map<String, dynamic>? params, Map<String, String>? header}) async {
    String respJson = await compute<Map<String, dynamic>, String>((data) {
      String dataStr = json.encode(data);
      Pointer<Int8> resp = goGet(GoString.fromString(dataStr));
      String respJson = resp.cast<Utf8>().toDartString();
      return respJson;
    }, {
      'url': url,
      'params': params,
      'header': header,
    });
    log('响应' + respJson);
    return toResponseData(respJson);
  }

  // 转换返回值为ResponseData
  toResponseData(String data){
    Map<String, dynamic> map = json.decode(data.toString());
    ResponseData responseData = ResponseData.fromJson(map);
    return responseData;
  }

}
