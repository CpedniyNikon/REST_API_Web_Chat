import 'package:flutter/material.dart';

class MainChat extends StatefulWidget {
  MainChat({Key? key}) : super(key: key);

  @override
  State<MainChat> createState() => _MainChatState();
}

class _MainChatState extends State<MainChat> {
  List<String> _notificationEmails = [
    "fdf","fdf"
  ];

  final TextEditingController _controller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return  Scaffold(
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
      onChanged: (String value) {
        if (value.substring(value.length - 1) == ',') {
          print('>>>>>> value = $value : controller = ${_controller.hashCode}');
          setState(() {
            _notificationEmails.add(value.substring(0, value.length - 1));
          });
          Future<void>.delayed(
            const Duration(milliseconds: 10),
            _controller.clear,
          );
          print(_notificationEmails);
        }
      },
    );
  }
}
