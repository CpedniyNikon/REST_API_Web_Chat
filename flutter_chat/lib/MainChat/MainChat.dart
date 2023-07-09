import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_chat/MainChat/requests/get_messages.dart';
import 'package:flutter_chat/MainChat/requests/write.dart';

class MainChat extends StatefulWidget {
  MainChat({Key? key}) : super(key: key);

  @override
  State<MainChat> createState() => _MainChatState();
}

class _MainChatState extends State<MainChat> {
  final List<String> _notificationEmails = [];
  late Timer timer;

  void update() {
    //this is a new Function
    setState(() {});
  }

  @override
  void initState() {
    timer = Timer.periodic(
        const Duration(
          seconds: 1,
        ),
        (Timer t) => getMessages(_notificationEmails, update));
    super.initState();
  }

  @override
  void dispose() {
    timer.cancel();
    super.dispose();
  }

  final TextEditingController _controller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          children: [
            _buildNotificationEmailsInput(),
            Expanded(
              child: ListView.builder(
                itemCount: _notificationEmails.length,
                itemBuilder: (_, int idx) => ListTile(
                  title: Text(_notificationEmails[idx]),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildNotificationEmailsInput() {
    return TextFormField(
      controller: _controller,
      validator: (String? value) {
        debugPrint('VALIDATOR: $value');
        if (value!.isEmpty) {
          return 'Emails Required';
        }
        return null;
      },
      onFieldSubmitted: write,
    );
  }
}
