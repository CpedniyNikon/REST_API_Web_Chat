import 'dart:convert';

import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_chat/Routes.dart';
import 'package:flutter_chat/utils/UserData.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;

class LoginPage extends StatefulWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
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
            color: Color.fromRGBO(24, 24, 32, 1),
          ),
          child: Center(
            child: Column(
              children: [
                const SizedBox(
                  height: 100,
                ),
                RichText(
                    text: const TextSpan(
                  text: "Sign in.",
                  style: TextStyle(fontSize: 60, color: Colors.white),
                )),
                const SizedBox(
                  height: 20,
                ),
                const SizedBox(
                  height: 10,
                ),
                ConstrainedBox(
                  constraints: const BoxConstraints(maxWidth: 400),
                  child: TextField(
                    decoration: InputDecoration(
                        contentPadding: const EdgeInsets.all(27),
                        enabledBorder: OutlineInputBorder(
                          borderSide: const BorderSide(
                            color: Colors.grey,
                            width: 1,
                          ),
                          borderRadius: BorderRadius.circular(20),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderSide: const BorderSide(
                            color: Colors.green,
                            width: 3,
                          ),
                          borderRadius: BorderRadius.circular(20),
                        ),
                        hintText: "login",
                        hintStyle: const TextStyle(color: Colors.grey)),
                    style: const TextStyle(color: Colors.white),
                    controller: loginController,
                  ),
                ),
                const SizedBox(
                  height: 20,
                ),
                ConstrainedBox(
                  constraints: const BoxConstraints(maxWidth: 400),
                  child: TextField(
                    decoration: InputDecoration(
                        contentPadding: const EdgeInsets.all(27),
                        enabledBorder: OutlineInputBorder(
                          borderSide: const BorderSide(
                            color: Colors.grey,
                            width: 1,
                          ),
                          borderRadius: BorderRadius.circular(20),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderSide: const BorderSide(
                            color: Colors.green,
                            width: 3,
                          ),
                          borderRadius: BorderRadius.circular(20),
                        ),
                        hintText: "password",
                        hintStyle: const TextStyle(color: Colors.grey)),
                    style: const TextStyle(color: Colors.white),
                    controller: passwordController,
                  ),
                ),
                const SizedBox(
                  height: 20,
                ),
                Container(
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(20),
                      gradient: const LinearGradient(colors: [
                        Color.fromRGBO(187, 63, 221, 1),
                        Color.fromRGBO(251, 109, 169, 1),
                        Color.fromRGBO(255, 159, 124, 1),
                      ], begin: Alignment.topLeft, end: Alignment.bottomRight)),
                  child: ElevatedButton(
                    onPressed: () =>
                        login(loginController.text, passwordController.text),
                    style: ElevatedButton.styleFrom(
                      fixedSize: const Size(400, 55),
                      backgroundColor: Colors.transparent,
                      shadowColor: Colors.transparent,
                    ),
                    child: const Text(
                      "Log in",
                      style: TextStyle(
                        color: Colors.black,
                        fontWeight: FontWeight.w600,
                        fontSize: 20,
                      ),
                    ),
                  ),
                ),
                const SizedBox(
                  height: 10,
                ),
                RichText(
                  text: TextSpan(
                    text: "No account?",
                    style: const TextStyle(color: Colors.white),
                    children: [
                      TextSpan(
                          text: " Register Now! ",
                          style: const TextStyle(color: Colors.red),
                          recognizer: TapGestureRecognizer()
                            ..onTap = () {
                              debugPrint("tapped");
                              Get.rootDelegate.toNamed(Routes.register);
                            }),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

Future<void> login(String login, String password) async {
  debugPrint("login");
  var url = 'http://localhost:8080/auth/sign-in';
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
    var output = json.decode(response.body);
    var text = output["text"];
    debugPrint(text);
    if (text == "u just logged in") {
      Get.rootDelegate.toNamed(Routes.chat);
      UserData.id = output["id"];
      debugPrint("id := ${UserData.id}");
    }
  });
}
