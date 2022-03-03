// post 请求
function post(url, params, header, contentType, isEncrypt) {
  var resp = global.Post(url, params, header, contentType, isEncrypt);
  if (!resp) {
    return {
      'err': '未知错误',
      'err_code': -1
    };
  }
  return JSON.parse(resp);
}

// get 请求
function get(url, params, header, contentType, isEncrypt) {
  var resp = global.Get(url, params, header, contentType, isEncrypt);
  if (!resp) {
    return {
      'err': '未知错误',
      'err_code': -1
    };
  }
  return JSON.parse(resp);
}

// put 请求
function put(url, params, header, contentType, isEncrypt) {
  var resp = global.Put(url, params, header, contentType, isEncrypt);
  if (!resp) {
    return {
      'err': '未知错误',
      'err_code': -1
    };
  }
  return JSON.parse(resp);
}

// delete 方法
function del(url, params, header, contentType, isEncrypt) {
  var resp = global.Delete(url, params, header, contentType, isEncrypt);
  if (!resp) {
    return {
      'err': '未知错误',
      'err_code': -1
    };
  }
  return JSON.parse(resp);
}

// request 请求
function request(method, url, params, header, contentType, isEncrypt) {
  var resp = global.Delete(url, params, header, contentType, isEncrypt);
  if (!resp) {
    return {
      'err': '未知错误',
      'err_code': -1
    };
  }
  return JSON.parse(resp);
}

// 获取版本信息
function getVersion() {
  var resp = global.GetVersion();
  return JSON.parse(resp);
}

// 设置请求根地址
function setBaseAddress(baseAddress) {
  var resp = global.SetBaseAddress(baseAddress);
  if (!resp) {
    return true;
  }
  console.log('设置根地址错误', JSON.parse(resp));
  return false;
}

// 设置超时 秒
function setTimeout(timeout) {
  var resp = global.SetTimeout(timeout);
  if (!resp) {
    return true;
  }
  console.log('设置根地址错误', JSON.parse(resp));
  return false;
}

// 设置公共请求头
function setHeader(name, value) {
  var resp = global.SetHeader(name, value);
  if (!resp) {
    return true;
  }
  console.log('设置根地址错误', JSON.parse(resp));
  return false;
}

// 导出
module.exports = {
  post: post,
  get: get,
  put: put,
  delete: del,
  request: request,
  getVersion: getVersion,
  setBaseAddress: setBaseAddress,
  setTimeout: setTimeout,
  setHeader: setHeader,
}
