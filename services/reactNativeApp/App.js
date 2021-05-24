/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import 'react-native-gesture-handler';
import React from 'react';
import {NavigationContainer} from '@react-navigation/native';
import {createStackNavigator} from '@react-navigation/stack';

import SignInScreen from './src/screens/SignInScreen';
import SignupScreen from './src/screens/SignupScreen';

const Stack = createStackNavigator();
const UserContext = React.createContext(null);

class App extends React.Component {
  state = {
    /** @type {IUser|null} */ user: null,
  };

  /**
   *
   * @param {IUser|null} user
   */
  setUser(user) {
    UserContext.this.setState({
      ...this.state,
      user: user,
    });
  }

  render() {
    return (
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen
            name="SignIn"
            options={{
              title: 'Hybride Entwicklung React Native',
            }}
            component={SignInScreen}
            initialParams={{setUser: this.setUser}}
          />
          <Stack.Screen
            name="Signup"
            options={{title: 'Hybride Entwicklung React Native'}}
            component={SignupScreen}
          />
          <Stack.Screen
            name="SignedIn"
            options={{title: 'Hybride Entwicklung React Native'}}
            component={SignInScreen}
          />
        </Stack.Navigator>
      </NavigationContainer>
    );
  }
}

export default App;
