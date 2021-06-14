import React, {Component} from 'react';
import {Button, StyleSheet, TextInput, View} from 'react-native';
import apiConf from '../config/api';
import ToastAndroid from 'react-native/Libraries/Components/ToastAndroid/ToastAndroid';

export default class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: '',
      password: '',
      passwordRepeat: '',
    };
  }

  onSignup() {
    const {username, password, passwordRepeat} = this.state;

    if (username.length === 0) {
      this.showToast('Username darf nicht leer sein');
      return;
    }

    if (passwordRepeat !== password) {
      this.showToast('Passwörter stimmen nicht überein');
      return;
    }

    if (password.length === 0) {
      this.showToast('Passwort darf nicht leer sein');
      return;
    }

    fetch(`${apiConf.apiUrl}/signup`, {
      method: 'POST',
      body: JSON.stringify({
        username,
        password,
      }),
    })
      .then(response => {
        const statusCode = response.status;
        const data = response.json();
        return Promise.all([statusCode, data]);
      })
      .then(([statusCode, data]) => {
        console.log({statusCode, data});
        if (statusCode == 201) {
          this.showToast('Registrierung erfolgreich');
        } else {
          this.showToast('Etwas ist schief gelaufen');
        }
      })
      .catch(console.error);
  }

  showToast(text) {
    ToastAndroid.showWithGravity(text, ToastAndroid.SHORT, ToastAndroid.BOTTOM);
  }

  render() {
    return (
      <View style={styles.container}>
        <TextInput
          value={this.state.username}
          onChangeText={username => this.setState({username})}
          placeholder={'Username'}
          style={styles.input}
        />
        <TextInput
          value={this.state.password}
          onChangeText={password => this.setState({password})}
          placeholder={'Passwort'}
          secureTextEntry={true}
          style={styles.input}
        />
        <TextInput
          value={this.state.passwordRepeat}
          onChangeText={passwordRepeat => this.setState({passwordRepeat})}
          placeholder={'Passwort wiederholen'}
          secureTextEntry={true}
          style={styles.input}
        />

        <Button
          title={'Registrieren'}
          style={styles.input}
          onPress={this.onSignup.bind(this)}
        />
        <Button
          title={'Zum Login'}
          style={styles.input}
          onPress={() => this.props.navigation.navigate('SignIn')}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: '#ecf0f1',
  },
  input: {
    width: 200,
    height: 44,
    padding: 10,
    borderWidth: 1,
    borderColor: 'black',
    marginBottom: 10,
  },
});
