import 'package:flutter/material.dart';
import 'package:flutter_app/signInPage.dart';
import 'package:flutter_app/signupPage.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'Hybride Entwicklung Flutter'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final pageView = PageView(
    controller: PageController(
      initialPage: 0,
    ),
    children: [
      SignInPage(),
      SignUpPage(),
    ],
  );

  @override
  Widget build(BuildContext context) {
    AppBar topHeader = AppBar(
      title: Text(widget.title),
    );

    return Scaffold(
      appBar: topHeader,
      body: pageView,
    );
  }
}
