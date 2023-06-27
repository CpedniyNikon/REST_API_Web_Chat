import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_chat/Routes.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;

class Authorization extends StatelessWidget {
  Authorization({Key? key}) : super(key: key);
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
                padding: const EdgeInsets.only(left: 100, right: 100),
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
                          onPressed: () => login(
                              loginController.text, passwordController.text),
                          icon: Image.asset('images/image.png')),
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

Future<void> login(String login, String password) async {

  var url = 'http://localhost:8080/auth/sign-up';
  await http
      .post(Uri.parse(url),
      headers: {},
      body: json.encode({
        "login": login,
        "password": password,
      }))
      .then((http.Response response) {
    debugPrint("Response status: ${response.statusCode}");
    debugPrint("Response body: ${response.contentLength}");
    debugPrint(response.body);
  });
  // Get.rootDelegate.toNamed(Routes.chat);
}
