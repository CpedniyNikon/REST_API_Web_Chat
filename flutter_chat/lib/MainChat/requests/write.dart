import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_chat/utils/UserData.dart';
import 'package:http/http.dart' as http;

write(String message) async {

  debugPrint("write");
  debugPrint(message);
  var url = 'http://localhost:8081/chat/write';
  await http
      .post(Uri.parse(url),
          headers: {},
          body: json.encode({
            "message": message,
            "userId": UserData.id,
          }))
      .then((http.Response response) {
    debugPrint("Response status: ${response.statusCode}");
    debugPrint("Response body: ${response.contentLength}");
    debugPrint(response.body);
  });
}
