import 'dart:convert';
import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_app/environment.dart';
import 'package:flutter_app/signedInPage.dart';
import 'package:flutter_app/user.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:http/http.dart' as http;

class SignInPage extends StatefulWidget {
  @override
  _State createState() => _State();
}

class _State extends State<SignInPage> {
  final usernameController = TextEditingController(text: "");
  final passwordController = TextEditingController(text: "");
  var isSigningIn = false;

  onSignIn() async {
    setState(() => isSigningIn = true);

    try {
      final response = await apiSignIn(
          username: usernameController.text, password: passwordController.text
      );

      // simulate delay that would be present if the api was web based
      await Future.delayed(Duration(seconds: 1));

      if (response.statusCode == 200) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text("Login Erfolgreich"),
          ),
        );

        var user = User.fromJson(jsonDecode(response.body));

        Navigator.push(
          context,
          MaterialPageRoute(
            builder: (context) => SignedInPage(
              user: user,
            ),
          ),
        );
      } else if (response.statusCode == 401) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text("Login Ungültig"),
          ),
        );
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text("Etwas ist schief gelaufen!"),
          ),
        );
      }
    } catch (err) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text("Server nicht erreichbar"),
        ),
      );
      print(err);
    }

    setState(() => isSigningIn = false);
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          TextField(
            controller: usernameController,
            decoration: InputDecoration(
              border: OutlineInputBorder(),
              hintText: "Username",
            ),
          ),
          TextField(
            controller: passwordController,
            autofillHints: [
              AutofillHints.password,
            ],
            obscureText: true,
            decoration: InputDecoration(
              border: OutlineInputBorder(),
              hintText: "Password",
            ),
          ),
          SizedBox(
            width: double.infinity,
            child: ElevatedButton(
              onPressed: isSigningIn ? null : onSignIn,
              child: isSigningIn ? SpinKitRing(
                      color: Colors.white,
                      size: 20.0,
                      lineWidth: 3,
                    ) : Text("Einloggen"),
            ),
          ),
        ],
      ),
    );
  }
}

Future<http.Response> apiSignIn({username, password}) {
  return http.post(
    Uri.http(EnvironmentConfig.ApiUrl, 'signIn'),
    headers: <String, String>{
      'Content-Type': 'application/json',
    },
    body: jsonEncode(
        <String, String>{'username': username, 'password': password}),
  );
}
