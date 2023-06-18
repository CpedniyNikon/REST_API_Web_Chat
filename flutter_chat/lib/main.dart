import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:url_strategy/url_strategy.dart';

void main() {
  setPathUrlStrategy();
  runApp(MaterialApp(
    theme: ThemeData(primaryColor: Colors.purple),
    debugShowCheckedModeBanner: false,
    home: Home(),
  ));
}

class Home extends StatelessWidget {
  TextEditingController loginController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Container(
          height: MediaQuery.of(context).size.height,
          width: MediaQuery.of(context).size.width,
          decoration: const BoxDecoration(
            gradient: LinearGradient(
              colors: [
                Color(0xFF8A2387),
                Color(0xFFE94057),
                Color(0xFFF27121),
              ],
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
            ),
          ),
          child: ListView(
            children: [
              const SizedBox(
                height: 120,
              ),
              Container(
                padding: const EdgeInsets.only(left: 20, right: 20),
                child: Container(
                  decoration: const BoxDecoration(
                      color: Colors.white,
                      borderRadius: BorderRadius.all(Radius.circular(20))),
                  child: Column(
                    children: [
                      const SizedBox(
                        height: 20,
                      ),
                      Container(
                        padding: const EdgeInsets.only(left: 20, right: 20),
                        child: TextField(
                          style: const TextStyle(color: Colors.black),
                          controller: loginController,
                          decoration: const InputDecoration(
                            border: InputBorder.none,
                            hintText: 'login',
                          ),
                        ),
                      ),
                      Container(
                        padding: const EdgeInsets.only(left: 20, right: 20),
                        child: TextField(
                          style: const TextStyle(color: Colors.black),
                          controller: passwordController,
                          decoration: const InputDecoration(
                            border: InputBorder.none,
                            hintText: 'password',
                          ),
                        ),
                      ),
                      IconButton(
                        onPressed: () => login(),
                        icon: const Icon(
                          Icons.ice_skating,
                        ),
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

Future<void> login() async {
  debugPrint('request');
  var url = 'http://localhost:8080/auth';
  await http
      .post(Uri.parse(url),
      headers: {
        "Content-Type": "application/json",
      },
      body: json.encode({}))
      .then((http.Response response) {
    debugPrint("Response status: ${response.statusCode}");
    debugPrint("Response body: ${response.contentLength}");
    debugPrint(response.body);
  });
}
