import 'package:flutter_test/flutter_test.dart';

import 'package:gohttp/gohttp.dart';

void main() {
  test('adds one to input values', () {
    final http = Http();
    expect(http.post('/v1/demo'), '');
  });
}
