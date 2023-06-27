import 'package:flutter_chat/AppPages.dart';
import 'package:flutter_chat/RouterDelegate.dart';
import 'package:get/get_navigation/src/root/get_material_app.dart';
import 'package:get/get_navigation/src/routes/transitions_type.dart';
import 'package:url_strategy/url_strategy.dart';
import 'package:flutter/material.dart';

void main() {
  setPathUrlStrategy();
  runApp(GetMaterialApp.router(
    debugShowCheckedModeBanner: false,
    defaultTransition: Transition.fade,
    getPages: AppPages.pages,
    routerDelegate: AppRouterDelegate(),
  ));
}
