import React, {Component} from 'react';
import {Button, Modal, StyleSheet, Text, TextInput, View} from 'react-native';
import apiConf from '../config/api';
import ToastAndroid from 'react-native/Libraries/Components/ToastAndroid/ToastAndroid';
import Pressable from 'react-native/Libraries/Components/Pressable/Pressable';

/**
 * @typedef IUser
 * @type {Object}
 * @property {string} id
 * @property {string} username
 */

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: '',
      password: '',
      /** @type {IUser|null} */ user: null,
    };
  }

  setUser(user) {
    this.setState({...this.state, user: user});
  }

  onLogin() {
    const {username, password} = this.state;

    fetch(`${apiConf.apiUrl}/signIn`, {
      method: 'POST',
      body: JSON.stringify({
        username,
        password,
      }),
    })
      .then(response => {
        console.log(response);
        const statusCode = response.status;
        const data = statusCode === 200 ? response.json() : null;
        return Promise.all([statusCode, data]);
      })
      .then(([statusCode, data]) => {
        console.log({statusCode, data, props: this.props.setUser});
        this.setUser(data);
        if (statusCode == 200) {
          this.showToast('Login erfolgreich');
        } else {
          this.showToast('Login ungültig');
        }
      })
      .catch(console.error);
  }

  showToast(text) {
    ToastAndroid.showWithGravity(text, ToastAndroid.SHORT, ToastAndroid.BOTTOM);
  }

  render() {
    return this.state.user ? (
      <View style={styles.centeredView}>
        <Modal
          animationType="slide"
          // transparent={true}
          // visible={this.state.user}
          onRequestClose={() => {
            this.setUser(null);
          }}>
          <View style={styles.centeredView}>
            <View style={styles.modalView}>
              <Text style={styles.modalTitle}>
                Willkommen {this.state.user.username}
              </Text>
              <Text style={styles.modalText}>Id: {this.state.user.id}</Text>
              <Text style={styles.modalText}>
                Username: {this.state.user.username}
              </Text>
              <Pressable
                style={styles.button}
                onPress={() => this.setUser(null)}>
                <Text style={styles.textStyle}>Zurück zum Login</Text>
              </Pressable>
            </View>
          </View>
        </Modal>
      </View>
    ) : (
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
          placeholder={'Password'}
          secureTextEntry={true}
          style={styles.input}
        />

        <Button
          title={'Login'}
          onPress={this.onLogin.bind(this)}
          style={styles.button}
        />
        <Button
          title={'Zum Registrieren'}
          onPress={() => this.props.navigation.navigate('Signup')}
          style={styles.button}
        />
      </View>
    );
  }
}

const styles = StyleSheet.create({
  centeredView: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 22,
  },
  modalView: {
    margin: 20,
    backgroundColor: 'white',
    borderRadius: 20,
    padding: 35,
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  button: {
    borderRadius: 20,
    padding: 10,
    elevation: 2,
    marginBottom: 20,
  },
  textStyle: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  modalTitle: {
    marginBottom: 15,
    fontSize: 24,
    textAlign: 'center',
  },
  modalText: {
    marginBottom: 15,
    textAlign: 'center',
  },
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
    borderRadius: 5,
    borderColor: 'black',
    marginBottom: 15,
  },
});
