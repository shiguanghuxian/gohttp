// index.js
// 获取应用实例
const app = getApp()

require('./assets/wasm_exec.js');
var http = require('./assets/http.js');

Page({
  data: {
    userInfo: {},
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs'
    })
  },

  async onReady() {
    console.log('onReady');
    // 在小程序基础类库的global对象上，增加console对象。
    global.console = console
    // 使用小程序类库的WXWebAssembly，初始化Go运行环境。
    await this.initGo()
  },

  onLoad() {
    console.log('onLoad');
  },

  async initGo() {
    const go = new global.Go();
    try {
      const result = await WXWebAssembly.instantiate('/pages/index/assets/libgohttp.wasm', go.importObject);
      var msg = 'Go初始化成功,在小程序调试窗口查看console的信息。';
      console.log('initGo99', msg);

      // 运行go程序的main()方法
      go.run(result.instance);
      // 设置基础信息
      http.setBaseAddress('https://localhost:9999'); // 需https安全域名
      http.setTimeout(10);
      http.setHeader("ZXP-DEMO", "test_header");
    } catch (err) {
      console.error('initGo', err);
    }
  },

  // 获取版本信息
  async bindVersion() {
    const version = http.getVersion();
    console.log(version);
    wx.showModal({
      content: `版本：${version.version}
GO版本：${version.go_version}
git hash：${version.git_hash}
构建时间：${version.build_time}`,
      showCancel: false,
      title: 'gohttp版本信息'
    })
  },


  async bindPost() {
    const params = {
      "a": 99,
      "b": "abc"
    };
    var resp = http.post('/v1/post', JSON.stringify(params), null, "application/json; charset=utf-8", true);
    console.log(resp);
    if (resp.err_code != 0) {
      wx.showModal({
        content: `err: ${resp.err} err_code: ${resp.err_code}`,
        showCancel: false,
      })
      return;
    }
    wx.showModal({
      content: `code: ${resp.status_code} status: ${resp.status}`,
      showCancel: false,
    })
  },

  async bindGET() {
    const params = {
      "a": 99,
      "b": "100f4d4aa98f70d821e58a9c5e813b87100f4d4aa98f70d821e58a9c5e813b87100f4d4aa98f70d821e58a9c5e813b87"
    };
    var resp = await http.get('/v1/get', JSON.stringify(params), null, "", true);
    console.log(resp);
    if (resp.err_code != 0) {
      wx.showModal({
        content: `err: ${resp.err} err_code: ${resp.err_code}`,
        showCancel: false,
      })
      return;
    }
    wx.showModal({
      content: `code: ${resp.status_code} status: ${resp.status}`,
      showCancel: false,
    })
  },

  async bindPUT() {
    const params = {
      "a": 99,
      "b": "abc"
    };
    var resp = await http.put('/v1/put', JSON.stringify(params), null, "application/x-www-form-urlencoded", true);
    console.log(resp);
    if (resp.err_code != 0) {
      wx.showModal({
        content: `err: ${resp.err} err_code: ${resp.err_code}`,
        showCancel: false,
      })
      return;
    }
    wx.showModal({
      content: `code: ${resp.status_code} status: ${resp.status}`,
      showCancel: false,
    })
  },

  async bindDELETE() {
    const params = {
      "a": 99,
      "b": "abc"
    };
    var resp = await http.delete('/v1/delete', JSON.stringify(params), null, '', false);
    console.log(resp);
    if (resp.err_code != 0) {
      wx.showModal({
        content: `err: ${resp.err} err_code: ${resp.err_code}`,
        showCancel: false,
      })
      return;
    }
    wx.showModal({
      content: `code: ${resp.status_code} status: ${resp.status}`,
      showCancel: false,
    })
  },


})