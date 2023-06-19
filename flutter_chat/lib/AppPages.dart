import 'package:flutter_chat/Routes.dart';
import 'package:flutter_chat/auth/Authorization.dart';
import 'package:get/get_navigation/src/routes/get_route.dart';

abstract class AppPages {
  static final pages = [
    GetPage(
      name: Routes.authentication,
      page: () => Authorization(),
    ),
  ];
}