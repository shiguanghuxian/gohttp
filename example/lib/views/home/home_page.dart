import 'dart:developer';
import 'dart:io';

import 'package:example/http/http.dart';
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:gohttp/gohttp.dart';
import 'package:path_provider/path_provider.dart';

class HomePage extends StatefulWidget {
  const HomePage({
    Key? key,
  }) : super(key: key);
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  void initState() {
    super.initState();
    _setCookiePath();
  }

  @override
  void dispose() {
    super.dispose();
  }

  // 设置请求根路径
  _setCookiePath() async {
    Directory tempDir = await getTemporaryDirectory();
    if (tempDir != null) {
      Directory cacheCookie = Directory(tempDir.path + "/" + "cache_cookie");
      if (!cacheCookie.existsSync()) {
        cacheCookie.create();
      }
      publicHttp
          .setCookiePath(cacheCookie.path);
    }
  }

  // 显示版本信息
  void _showVersion() async {
    // 输出版本信息
    dynamic version = publicHttp.getVersion();
    log('gohttp版本: ${version.version}');

    String dateStr =
        DateTime.fromMillisecondsSinceEpoch(int.parse(version.buildTime) * 1000)
            .toString();
    EasyLoading.showInfo(
        "版本：${version.version}\nGO版本：${version.goVersion}\ngit hash：${version.gitHash}\n构建时间：$dateStr");
  }

  void _post() async {
    ResponseData respJson = await publicHttp.post(
      '/v1/post',
      params: {
        "a": 1,
        "b": "bbb",
      },
      contentType: 'application/json; charset=utf-8',
      encrypt: true,
    );
    EasyLoading.showInfo(
        "code: ${respJson.statusCode} \n status: ${respJson.status} \n ${respJson.body}");
  }

  void _get() async {
    ResponseData respJson = await publicHttp.get(
      '/v1/get',
      params: {
        "a": 2,
        "b": "bbb",
      },
      encrypt: true,
    );
    EasyLoading.showInfo(
        "code: ${respJson.statusCode} \n status: ${respJson.status} \n ${respJson.body}");
  }

  void _put() async {
    ResponseData respJson = await publicHttp.put('/v1/put', params: {
      "a": 3,
      "b": "bbb",
    });
    EasyLoading.showInfo(
        "code: ${respJson.statusCode} \n status: ${respJson.status} \n ${respJson.body}");
  }

  void _delete() async {
    ResponseData respJson = await publicHttp.delete('/v1/delete', params: {
      "a": 5,
      "b": "bbb",
    });
    EasyLoading.showInfo(
        "code: ${respJson.statusCode} \n status: ${respJson.status} \n ${respJson.body}");
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        title: const Text('gohttp'),
      ),
      body: SafeArea(
        child: ListView(
          children: [
            const SizedBox(
              height: 20,
            ),
            TextButton(
              onPressed: () {
                _showVersion();
              },
              child: const Text('Version'),
            ),
            const SizedBox(
              height: 20,
            ),
            TextButton(
              onPressed: () {
                _post();
              },
              child: const Text('POST'),
            ),
            const SizedBox(
              height: 20,
            ),
            TextButton(
              onPressed: () {
                _get();
              },
              child: const Text('GET'),
            ),
            const SizedBox(
              height: 20,
            ),
            TextButton(
              onPressed: () {
                _put();
              },
              child: const Text('PUT'),
            ),
            const SizedBox(
              height: 20,
            ),
            TextButton(
              onPressed: () {
                _delete();
              },
              child: const Text('DELETE'),
            ),
          ],
        ),
      ),
    );
  }
}
