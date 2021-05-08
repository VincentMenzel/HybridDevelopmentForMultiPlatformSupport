class EnvironmentConfig {
  static const ApiUrl = String.fromEnvironment('api_url', defaultValue: "192.168.5.129:8080");
}