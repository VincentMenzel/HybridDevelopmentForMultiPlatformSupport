import 'dart:convert';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:http/http.dart' as http;

import 'environment.dart';

class SignUpPage extends StatefulWidget {
  const SignUpPage({Key key}) : super(key: key);

  @override
  _SignUpPageState createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  var isSigningUp = false;
  var usernameController = TextEditingController(text: "");
  var passwordController = TextEditingController(text: "");
  var passwordRepeatController = TextEditingController(text: "");

  onSignUp() async {
    var password = passwordController.text;
    var passwordRepeat = passwordRepeatController.text;
    var username = usernameController.text.trim();

    if (username.length == 0) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text("Username darf nicht leer sein"),
          duration: Duration(seconds: 1),
        ),
      );
      return;
    }

    if (passwordRepeat != password) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text("Passwörter stimmen nicht überein"),
          duration: Duration(seconds: 1),
        ),
      );
      return;
    }

    if (password.length == 0) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text("Passwort darf nicht leer sein"),
          duration: Duration(seconds: 1),
        ),
      );
      return;
    }

    setState(() => isSigningUp = true);

    try {
      // simulate delay that would be present if the api was web based
      await Future.delayed(Duration(seconds: 1));

      var res = await apiSignup(username: username, password: password);

      if (res.statusCode == 201) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text("Registrierung erfolgreich"),
            duration: Duration(seconds: 1),
          ),
        );
      } else {
        print(res.body);
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text("Etwas ist schief gelaufen"),
            duration: Duration(seconds: 1),
          ),
        );
      }
    } catch (err) {
      print(err);
    }

    setState(() => isSigningUp = false);
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          TextField(
            controller: usernameController,
            autofillHints: [
              AutofillHints.newUsername,
            ],
            decoration: InputDecoration(
              border: OutlineInputBorder(),
              hintText: "Username",
            ),
          ),
          TextField(
            controller: passwordController,
            autofillHints: [
              AutofillHints.newPassword,
            ],
            decoration: InputDecoration(
              border: OutlineInputBorder(),
              hintText: "Passwort",
            ),
          ),
          TextField(
            controller: passwordRepeatController,
            autofillHints: [
              AutofillHints.newPassword,
            ],
            decoration: InputDecoration(
              border: OutlineInputBorder(),
              hintText: "Passwort wiederholen",
            ),
          ),
          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: isSigningUp ? null : onSignUp,
              child: isSigningUp
                  ? SpinKitRing(
                      color: Colors.white,
                      size: 20.0,
                      lineWidth: 3,
                    )
                  : Text("Registrieren"),
            ),
          ),
        ],
      ),
    );
  }
}

Future<http.Response> apiSignup({username, password}) {
  return http.post(
    Uri.http(EnvironmentConfig.ApiUrl, 'signup'),
    headers: <String, String>{
      'Content-Type': 'application/json',
    },
    body: jsonEncode(
        <String, String>{'username': username, 'password': password}),
  );
}
