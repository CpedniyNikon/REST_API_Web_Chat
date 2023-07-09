import 'package:flutter_chat/MainChat/MainChat.dart';
import 'package:flutter_chat/auth/LoginPage.dart';
import 'package:flutter_chat/auth/RegisterPage.dart';
import 'package:flutter_chat/routes/Routes.dart';
import 'package:get/get_navigation/src/routes/get_route.dart';

abstract class AppPages {
  static final pages = [
    GetPage(
      name: Routes.chat,
      page: () => MainChat(),
    ),
    GetPage(
      name: Routes.login,
      page: () => const LoginPage(),
    ),
    GetPage(
      name: Routes.register,
      page: () => const RegistrationPage(),
    ),
  ];
}
