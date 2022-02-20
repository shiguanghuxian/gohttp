part of gohttp;

class ResponseData {
  late String err;
  late int errCode;
  late String status;
  late int statusCode;
  late String body;

  ResponseData(
      this.err, this.errCode, this.status, this.statusCode, this.body);

  ResponseData.fromJson(Map<String, dynamic> json) {
    err = json['err'];
    errCode = json['err_code'];
    status = json['status'];
    statusCode = json['status_code'];
    body = json['body'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['err'] = err;
    data['err_code'] = errCode;
    data['status'] = status;
    data['status_code'] = statusCode;
    data['body'] = body;
    return data;
  }
}

class GoHttpVersion {
  String? buildTime;
  String? gitHash;
  String? goVersion;
  String? version;

  GoHttpVersion({this.buildTime, this.gitHash, this.goVersion, this.version});

  GoHttpVersion.fromJson(Map<String, dynamic> json) {
    buildTime = json['build_time'];
    gitHash = json['git_hash'];
    goVersion = json['go_version'];
    version = json['version'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = Map<String, dynamic>();
    data['build_time'] = buildTime;
    data['git_hash'] = gitHash;
    data['go_version'] = goVersion;
    data['version'] = version;
    return data;
  }
}

