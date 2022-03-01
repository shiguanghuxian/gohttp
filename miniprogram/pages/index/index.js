// index.js
// 获取应用实例
const app = getApp()

require('./assets/wasm_exec.js');

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
    // 在小程序基础类库的global对象上，增加console对象。
    global.console = console
    // 使用小程序类库的WXWebAssembly，初始化Go运行环境。
    await this.initGo()
  },

  async initGo() {
    var _that = this;
    const go = new global.Go();
    try {
      const result = await WXWebAssembly.instantiate('/pages/index/assets/libgohttp.wasm', go.importObject);
      var msg = 'Go初始化成功,在小程序调试窗口查看console的信息。';
      console.log('initGo99', msg);

      // 运行go程序的main()方法
      await go.run(result.instance);
      // 注意：在go程序的main()方法退出之前，小程序不会运行到这个位置。
      console.log('initGo', '运行完成');
    } catch (err) {
      console.error('initGo', err);
    }
  },

  onLoad() {

  },


  async bindPost() {
    const params = {
      "a": 99,
      "b": "abc"
    };
    var resp = await global.Post('/v1/post', JSON.stringify(params), null, "application/json; charset=utf-8", true);
    console.log(resp);
    if (resp.err_code != 0) {
      alert(`err: ${resp.err} err_code: ${resp.err_code}`);
      return;
    }
    alert(`code: ${resp.status_code} status: ${resp.status}`);
  },

  async get() {
    params = {
      "a": 99,
      "b": "100f4d4aa98f70d821e58a9c5e813b87100f4d4aa98f70d821e58a9c5e813b87100f4d4aa98f70d821e58a9c5e813b87"
    };
    var resp = await Get('/v1/get', JSON.stringify(params), null, "", true);
    console.log(resp);
    if (resp.err_code != 0) {
      alert(`err: ${resp.err} err_code: ${resp.err_code}`);
      return;
    }
    alert(`code: ${resp.status_code} status: ${resp.status}`);
  },

  async put() {
    params = {
      "a": 99,
      "b": "abc"
    };
    var resp = await Put('/v1/put', JSON.stringify(params), null, "application/x-www-form-urlencoded", true);
    console.log(resp);
    if (resp.err_code != 0) {
      alert(`err: ${resp.err} err_code: ${resp.err_code}`);
      return;
    }
    alert(`code: ${resp.status_code} status: ${resp.status}`);
  },

  async del() {
    params = {
      "a": 99,
      "b": "abc"
    };
    var resp = await Delete('/v1/delete', JSON.stringify(params));
    console.log(resp);
    if (resp.err_code != 0) {
      alert(`err: ${resp.err} err_code: ${resp.err_code}`);
      return;
    }
    alert(`code: ${resp.status_code} status: ${resp.status}`);
  },


})