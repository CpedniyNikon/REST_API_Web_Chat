import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_chat/utils/UserData.dart';
import 'package:http/http.dart' as http;

class Message {
  late String user;
  late String message;


  @override
  String toString() {
    return "$user $message";
  }
}

getMessages(List<String> messages, void Function() update) async {
  debugPrint("get_message");
  var url = 'http://localhost:8081/chat/get_messages';
  await http
      .post(Uri.parse(url), headers: {},
      body: json.encode({
        "userID": UserData.id
      }) )
      .then((http.Response response) {
    debugPrint("Response status: ${response.statusCode}");
    debugPrint("Response body: ${response.contentLength}");
    debugPrint(response.body);
    List<dynamic> parsedJson=  json.decode(response.body);

    messages.clear();
    for (var json in parsedJson) {
      Message message = Message();
      message.user = json['user'];
      message.message = json['message'];
      messages.add(message.toString());
    }
    update();
  });
}
