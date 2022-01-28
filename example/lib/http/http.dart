import 'package:gohttp/gohttp.dart';

// 全局公用http请求对象
final publicHttp = Http()
  ..setBaseAddress(HttpConfig.baseAddress)
  ..setTimeout(HttpConfig.timeout);

class HttpConfig {
  static const baseAddress = "https://www.baidu.com"; // 根请求地址
  static const timeout = 20; // 访问超时 15s
}

