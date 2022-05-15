# gohttp

使用go语言开发的客户端加密请求库，本代码提供代码实践，不建议直接用于生产，使用时至少将签名算法修改。

## 前言

需要一个抖音爬虫，去抓取一些数据，发现抖音很多接口都使用了动态连结库进行加密请求，与之相同的快手也做了类似方案。
从而感想到业务中或多或少都会遇到一些接口对于安全要求比较高，随产生开发一个加密请求库，以C动态链接库形式提供。
此库最佳的开发语言应为C或rust，但本人现阶段掌握的语言是go，go语言唯一缺点是打包动态库体积较大，经探索体积在5mb左右，
这个大小放在10年前不可容忍，但现阶段一个app体积一般都会冲破100mb，一些大型app接近1gb，所以10mb以内打包一个动态库完全可以忍受，
加上go语言夸平台的优点，此动态库可以方便的构建到各种cpu架构的各种系统下，结合优缺点go语言很合适做这件事。

经过评估go作为开发语言开发这样一个库可行。接下来就是库开发了需要测试，鉴于本人没有安卓和ios等原生开发能力，但拥有flutter开发能力，
并且flutter又同go语言一样是一个跨平台的方案，那么可以合理的理解为flutter使用此库可以在ios和安卓等平台运行，那么此库就是可用的。

## 代码结构

```
├── CHANGELOG.md
├── LICENSE
├── README.md
├── analysis_options.yaml
├── example // 示例文件夹
│   ├── README.md
│   ├── analysis_options.yaml
│   ├── android
│   ├── ios
│   ├── lib
│   │   ├── http
│   │   │   └── http.dart
│   │   ├── main.dart
│   │   └── views
│   │       └── home
│   │           └── home_page.dart
│   ├── macos
│   └── windows
├── go.mod
├── gosrc // gohttp核心go代码
│   ├── README.md
│   ├── export_c // c动态库
│   │   ├── Makefile
│   │   ├── clangwrap.sh
│   │   ├── file_cookie_jar.go // 本地cookie存储实现
│   │   ├── main.go // c导出函数
│   │   └── make.bat
│   ├── export_wasm // 浏览器标准wasm
│   │   ├── Makefile
│   │   └── main.go
│   ├── export_wx_wasm // 小程序wasm
│   │   ├── Makefile
│   │   └── main.go
│   ├── gohttp
│   │   ├── ascii
│   │   │   ├── print.go
│   │   │   └── print_test.go
│   │   ├── cookiejar
│   │   │   ├── cookie_jar.go
│   │   │   ├── default_cookie_jar.go
│   │   │   └── punycode.go
│   │   ├── errors.go
│   │   ├── funcs.go
│   │   ├── gohttp.go
│   │   ├── gohttplog
│   │   │   ├── env.go
│   │   │   ├── log.go
│   │   │   └── log_js.go
│   │   ├── model.go
│   │   └── request.go
│   ├── server // 服务端示例，使用gin，其它框架可参考中间件实现
│   │   ├── Makefile
│   │   ├── index.html
│   │   ├── main.go
│   │   └── static
│   │       └── wasm_exec.js
│   └── version
├── lib // flutter插件
│   ├── gohttp.dart
│   └── src
│       ├── godart.dart
│       ├── gohttp.dart
│       └── model.dart
├── miniprogram // 微信小程序
│   ├── app.js
│   ├── app.json
│   ├── app.wxss
│   ├── miniprogram_npm
│   │   └── text-encoder
│   │       ├── index.js
│   │       └── index.js.map
│   ├── package-lock.json
│   ├── package.json
│   ├── pages
│   │   ├── index
│   │   │   ├── assets
│   │   │   │   ├── http.js
│   │   │   │   └── wasm_exec.js // 经过改写的go wasm胶水js
│   │   │   ├── index.js
│   │   │   ├── index.json
│   │   │   ├── index.wxml
│   │   │   └── index.wxss
│   │   └── logs
│   │       ├── logs.js
│   │       ├── logs.json
│   │       ├── logs.wxml
│   │       └── logs.wxss
│   ├── project.config.json
│   ├── project.private.config.json
│   ├── sitemap.json
│   └── utils
│       └── util.js
├── pubspec.yaml

```

## 功能

1. 在原生端提供使用go导出的c动态库，可以支持ios、安卓、windows、macos、linux平台
2. 在web端导出wasm实现，可以支持浏览器、微信小程序。
3. 支持常见的get、post、put、delete和request自定义请求
4. c动态库支持cookie持久化存储
5. 支持设置请求根地址，如果请求地址以`http://`或`https://`开头，则不使用根地址
6. 支持设置公共请求头信息


## 备注

1. 签名算法`gosrc/gohttp/funcs.go:func Signature`
2. 微信小程序不能直接使用go提供的胶水文件`wasm_exec.js`
3. 使用`compute`函数解决flutter调用c函数UI主线程阻塞问题
