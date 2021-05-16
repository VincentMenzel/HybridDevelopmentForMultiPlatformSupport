import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/user.dart';

class SignedInPage extends StatefulWidget {
  final User user;

  const SignedInPage({this.user});

  @override
  _SignedInPageState createState() => _SignedInPageState();
}

class _SignedInPageState extends State<SignedInPage> {
  @override
  Widget build(BuildContext context) {
    AppBar topHeader = AppBar(
      title: Text('Willkommen ${widget.user.username}'),
    );

    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: topHeader,
        body: Container(
          decoration: BoxDecoration(),
          child: Center(
              child: Card(
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 0, vertical: 20),
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  ListTile(
                    leading: Icon(Icons.account_box),
                    title: Text("Eingeloggt als"),
                  ),
                  Column(
                    children: [
                      Text('Id: ${widget.user.id}'),
                      Text('Username: ${widget.user.username}'),
                    ],
                  ),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.end,
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      ElevatedButton(onPressed: () => Navigator.pop(context), child: Text("Zurück zum login"))
                    ],
                  )
                ],
              ),
            ),
          )
              // Column(
              //   mainAxisAlignment: MainAxisAlignment.center,
              //   children: [
              //     Text('Id: ${widget.user.id}'),
              //     Text('Username: ${widget.user.username}'),
              //   ],
              // ),
              ),
        ),
      ),
    );
  }
}
